package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/core/domain"
	"server/core/enums/userPrivilege"
	"server/core/ports"
	"server/tools/jwt"
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
}

func (h *AuthHandler) Register(ctx echo.Context) error {
	var user domain.CreateUser

	if err := ctx.Bind(&user); err != nil {
		return err
	}

	if err := ctx.Validate(&user); err != nil {
		return err
	}

	id, err := h.serviceUser.SaveUser(ctx.Request().Context(), &user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	jwtToken, err := jwt.CreateToken(id, userPrivilege.User.String())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.String(http.StatusOK, jwtToken)
}
