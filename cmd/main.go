package main

import (
	"github.com/labstack/echo/v4"
	"server/internal/server"
)

func main() {
	e := echo.New()
	serverInstance := server.GetServer()

	e.GET("/users/:id", serverInstance.ControllerUser.GetUser)
	e.GET("/users", serverInstance.ControllerUser.SaveUser)

	e.Logger.Fatal(e.Start(":1234"))
}
