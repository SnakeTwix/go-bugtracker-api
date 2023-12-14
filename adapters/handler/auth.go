package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/core/domain"
	"server/core/enums/cookies"
	"server/core/ports"
	"strconv"
	"time"
)

type AuthHandler struct {
	serviceUser ports.ServiceUser
}

var authHandler *AuthHandler

func GetAuthHandler(serviceUser ports.ServiceUser) *AuthHandler {
	if authHandler != nil {
		return authHandler
	}

	authHandler = &AuthHandler{
		serviceUser: serviceUser,
	}

	return authHandler
}

func (h *AuthHandler) RegisterRoutes(e *echo.Group) {
	e.POST("/auth/register", h.Register)
	e.POST("/auth/login", h.Login)
	e.POST("/auth/logout", h.Logout)
	e.POST("/auth/session", h.UpdateSession)
}

func getRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     cookies.RefreshToken,
		Value:    refreshToken,
		Path:     "/api/v1/auth/session",
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}
func getAccessCookie(accessToken string) *http.Cookie {
	return &http.Cookie{
		Name:     cookies.AccessToken,
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Minute * 30),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

// Register godoc
// @Summary      Registers a user
// @Description  Registers a user with default permissions
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param user body domain.CreateUser true "The user data to use when registering"
// @Success      200  {array}   domain.Token
// @Failure      400  {object}  error "Should only happen when there is already a user with the same username"
// @Failure      500  {object}  error "Shouldn't happen, but will if something fails"
// @Router       /api/v1/auth/register [POST]
func (h *AuthHandler) Register(ctx echo.Context) error {
	var user domain.CreateUser

	if err := ctx.Bind(&user); err != nil {
		return err
	}
	if err := ctx.Validate(&user); err != nil {
		return err
	}

	token, err := h.serviceUser.RegisterUser(ctx.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newRefreshCookie := getRefreshCookie(token.RefreshToken)
	newAccessToken := getAccessCookie(token.Jwt)

	ctx.SetCookie(newRefreshCookie)
	ctx.SetCookie(newAccessToken)

	return ctx.NoContent(http.StatusOK)
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	var user domain.LoginUser

	if err := ctx.Bind(&user); err != nil {
		return err
	}
	if err := ctx.Validate(&user); err != nil {
		return err
	}

	token, err := h.serviceUser.LoginUser(ctx.Request().Context(), &user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	newRefreshCookie := getRefreshCookie(token.RefreshToken)
	newAccessToken := getAccessCookie(token.Jwt)

	ctx.SetCookie(newRefreshCookie)
	ctx.SetCookie(newAccessToken)

	return ctx.NoContent(http.StatusOK)
}

func (h *AuthHandler) Logout(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	err = h.serviceUser.LogoutUser(ctx.Request().Context(), id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	newRefreshCookie := getRefreshCookie("")
	newAccessToken := getAccessCookie("")

	ctx.SetCookie(newRefreshCookie)
	ctx.SetCookie(newAccessToken)

	return ctx.NoContent(http.StatusOK)
}

func (h *AuthHandler) UpdateSession(ctx echo.Context) error {
	refreshCookie, err := ctx.Cookie(cookies.RefreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No refresh cookie provided")
	}
	token, err := h.serviceUser.UpdateSession(ctx.Request().Context(), &refreshCookie.Value)

	newRefreshCookie := getRefreshCookie(token.RefreshToken)
	newAccessToken := getAccessCookie(token.Jwt)

	ctx.SetCookie(newRefreshCookie)
	ctx.SetCookie(newAccessToken)

	return ctx.NoContent(http.StatusOK)
}
