package main

import (
	"github.com/labstack/gommon/log"
	"server/adapters/handler"
	"server/adapters/repository"
	"server/adapters/repository/migrations"
	"server/core/services"
	"server/tools/echo"
)

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

	e.Logger.Debug(e.Start(":1234"))
}
