package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/adapters/tools/jwt"
	"server/core/domain"
	"server/core/enums/cookies"
	"server/core/ports"
	"strconv"
)

type UserHandler struct {
	serviceUser ports.ServiceUser
}

var userHandler *UserHandler

func GetUserHandler(serviceUser ports.ServiceUser) *UserHandler {
	if userHandler != nil {
		return userHandler
	}

	userHandler = &UserHandler{
		serviceUser: serviceUser,
	}

	return userHandler
}

func (h *UserHandler) RegisterRoutes(e *echo.Group) {
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUser)
	e.GET("/user", h.GetCurrentUser)
}

// GetUser godoc
// @Summary      Get a user
// @Description  Get a user by id provided in the link
// @Tags         users
// @Produce      json
// @Param userId path uint64 true "The user id that you need to get"
// @Success      200  {object}   domain.GetUser
// @Failure      400  {object}  error "Happens if the id isn't numerical"
// @Failure      404  {object}  error "Happens if there is no user with such id"
// @Failure      500  {object}  error "Shouldn't happen, but will if something fails"
// @Router       /api/v1/users/{userId} [GET]
func (h *UserHandler) GetUser(ctx echo.Context) error {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		return echo.ErrBadRequest
	}

	user, err := h.serviceUser.GetUser(ctx.Request().Context(), id)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary      Get all users
// @Tags         users
// @Produce      json
// @Success      200  {array}   []domain.GetUser
// @Failure      500  {object}  error "Shouldn't happen, but will if something goes really wrong"
// @Router       /api/v1/users [GET]
func (h *UserHandler) GetUsers(ctx echo.Context) error {
	users, err := h.serviceUser.GetUsers(ctx.Request().Context())

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) SaveUser(ctx echo.Context) error {
	var user domain.CreateUser

	if err := ctx.Bind(&user); err != nil {
		return err
	}

	if err := ctx.Validate(&user); err != nil {
		return err
	}

	id, err := h.serviceUser.RegisterUser(ctx.Request().Context(), &user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, id)
}

func (h *UserHandler) GetCurrentUser(ctx echo.Context) error {
	accessCookie, err := ctx.Cookie(cookies.AccessToken)
	if err != nil {
		return echo.ErrUnauthorized
	}

	accessToken := accessCookie.Value
	tokenClaims, err := jwt.ParseUserClaims(accessToken)
	if err != nil {
		return err
	}

	user, err := h.serviceUser.GetUser(ctx.Request().Context(), tokenClaims.Id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}
