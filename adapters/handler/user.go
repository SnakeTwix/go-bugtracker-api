package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/core/domain"
	"server/core/ports"
	"strconv"
)

type UserHandler struct {
	serviceUser ports.ServiceUser
}

var httpHandler *UserHandler

func GetUserHandler(serviceUser ports.ServiceUser) *UserHandler {
	if httpHandler != nil {
		return httpHandler
	}

	httpHandler = &UserHandler{
		serviceUser: serviceUser,
	}

	return httpHandler
}

func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUser)
	e.POST("/users", h.SaveUser)
}

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

	err := h.serviceUser.SaveUser(ctx.Request().Context(), &user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
