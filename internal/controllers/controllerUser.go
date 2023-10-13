package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	ModelUser "server/internal/models"
	"server/internal/server"
)

var serverInstance = server.GetServer()

type createUserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type IControllerUser interface {
	SaveUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type ControllerUser struct{}

func (_ ControllerUser) SaveUser(c echo.Context) error {
	var createUser createUserDTO

	if err := c.Bind(&createUser); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	user := ModelUser.CreateUser{
		Name:     createUser.Name,
		Password: createUser.Password,
	}

	createdUser, httpErr := serverInstance.ModelUsers.CreateUser(user)

	if httpErr != nil {
		return httpErr
	}

	return c.JSON(http.StatusOK, createdUser)
}

type getUserDTO struct {
	Id int `param:"id"`
}

func (_ ControllerUser) GetUser(c echo.Context) error {
	var schema getUserDTO

	if err := c.Bind(&schema); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, err := serverInstance.ModelUsers.ReadUserById(schema.Id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (_ ControllerUser) UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "hi")
}

func (_ ControllerUser) DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "hi")
}
