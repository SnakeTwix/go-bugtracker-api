package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/core/domain"
	"server/core/ports"
	"server/tools/middleware"
	"strconv"
	"time"
)

type AuthHandler struct {
	serviceUser    ports.ServiceUser
	serviceSession ports.ServiceSession
}

var authHandler *AuthHandler

func GetAuthHandler(serviceUser ports.ServiceUser, serviceSession ports.ServiceSession) *AuthHandler {
	if authHandler != nil {
		return authHandler
	}

	authHandler = &AuthHandler{
		serviceUser:    serviceUser,
		serviceSession: serviceSession,
	}

	return authHandler
}

func (h *AuthHandler) RegisterRoutes(middleware *middleware.Middleware, group *echo.Group) {
	group.POST("/auth/register", h.Register)
	group.POST("/auth/login", h.Login)
	group.POST("/auth/logout", h.Logout)
	//e.POST("/auth/session", h.UpdateSession)
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

	createdUser, err := h.serviceUser.RegisterUser(ctx.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	IP := ctx.RealIP()
	session, err := h.serviceSession.CreateSession(ctx.Request().Context(), &IP, createdUser)
	if err != nil {
		return err
	}

	ctx.SetCookie(session.Cookie())

	return ctx.NoContent(http.StatusOK)
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	var loginUser domain.LoginUser

	if err := ctx.Bind(&loginUser); err != nil {
		return err
	}
	if err := ctx.Validate(&loginUser); err != nil {
		return err
	}

	user, err := h.serviceUser.LoginUser(ctx.Request().Context(), &loginUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	IP := ctx.RealIP()
	session, err := h.serviceSession.CreateSession(ctx.Request().Context(), &IP, user)
	if err != nil {
		return err
	}

	ctx.SetCookie(session.Cookie())

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

	logoutCookie := (&domain.Session{Expiry: time.Now()}).Cookie()
	ctx.SetCookie(logoutCookie)

	return ctx.NoContent(http.StatusOK)
}
