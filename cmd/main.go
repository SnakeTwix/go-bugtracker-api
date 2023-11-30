package main

import (
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"server/adapters/handler"
	"server/adapters/repository"
	"server/adapters/repository/migrations"
	"server/core/services"
	"server/tools/echo"

	_ "server/docs"
)

// @title Swagger Test
// @version 1.0
// @description My project test swagger
// @license.name Apache 2.0
//
// @host localhost:1234
// @BasePath /api/v1
func main() {
	e := echo.GetEchoInstance()
	apiV1 := e.Group("/api/v1")

	db := repository.InitDB()

	if err := migrations.RunMigrations(db); err != nil {
		log.Error(err)
		panic(err)
	}

	repoUser := repository.GetRepoUser(db)
	serviceUser := services.GetServiceUser(repoUser)
	userHandler := handler.GetUserHandler(serviceUser)
	authHandler := handler.GetAuthHandler(serviceUser)

	userHandler.RegisterRoutes(apiV1)
	authHandler.RegisterRoutes(apiV1)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Debug(e.Start(":1234"))
}
