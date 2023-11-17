package main

import (
	"github.com/labstack/gommon/log"
	"server/adapters/handler"
	"server/adapters/repository"
	"server/adapters/repository/migrations"
	"server/core/services"
	"server/initializer/echo"
)

func main() {
	e := echo.GetEchoInstance()
	db := repository.InitDB()

	if err := migrations.RunMigrations(db); err != nil {
		log.Error(err)
		panic(err)
	}

	repoUser := repository.GetRepoUser(db)
	serviceUser := services.GetServiceUser(repoUser)
	userHandler := handler.GetUserHandler(serviceUser)

	userHandler.RegisterRoutes(e)

	e.Logger.Debug(e.Start(":1234"))
}
