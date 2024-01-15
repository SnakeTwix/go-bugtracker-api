package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/core/domain"
	contextEnum "server/core/enums/context"
	"server/core/ports"
	"server/tools/middleware"
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

func (h *UserHandler) RegisterRoutes(middleware *middleware.Middleware, group *echo.Group) {
	group.GET("/users", h.GetUsers)
	group.GET("/users/:id", h.GetUser)
	group.GET("/users/current", h.GetCurrentUser)
	group.GET("/test", h.TestRoute, middleware.CheckLoggedIn)
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

//func (h *UserHandler) SaveUser(ctx echo.Context) error {
//	var user domain.CreateUser
//
//	if err := ctx.Bind(&user); err != nil {
//		return err
//	}
//
//	if err := ctx.Validate(&user); err != nil {
//		return err
//	}
//
//	id, err := h.serviceUser.RegisterUser(ctx.Request().Context(), &user)
//
//	if err != nil {
//		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//	}
//
//	return ctx.JSON(http.StatusOK, id)
//}

func (h *UserHandler) GetCurrentUser(ctx echo.Context) error {
	session, ok := ctx.Get(contextEnum.Session).(*domain.Session)
	if !ok {
		return echo.ErrUnauthorized
	}

	user, err := h.serviceUser.GetUser(ctx.Request().Context(), session.UserID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) TestRoute(ctx echo.Context) error {
	//fmt.Println(ctx.Get("session"))
	session, ok := ctx.Get("session").(*domain.Session)
	if !ok {
		fmt.Println("Isn't a session")
	}

	fmt.Print(session)

	return ctx.NoContent(200)
}
