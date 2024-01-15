package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"server/adapters/handler"
	"server/adapters/repository"
	"server/adapters/repository/migrations"
	"server/core/service"
	_ "server/docs"
	"server/tools/middleware"
	"server/tools/server"
)

// @title Swagger Test
// @version 1.0
// @description My project test swagger
// @license.name Apache 2.0
//
// @host localhost:1234
// @BasePath /api/v1
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := repository.InitDB()
	if err := migrations.RunMigrations(db); err != nil {
		log.Error(err)
		panic(err)
	}

	// Repos
	repoUser := repository.GetRepoUser(db)
	repoSession := repository.GetRepoSession(db)

	// Services
	serviceUser := service.GetServiceUser(repoUser)
	serviceSession := service.GetServiceSession(repoSession)
	serviceAuth := service.GetServiceAuth(repoUser, repoSession)

	services := server.Services{
		ServiceSession: serviceSession,
		ServiceUser:    serviceUser,
	}

	serverInstance := server.GetServerInstance(&services)
	apiV1 := serverInstance.Echo.Group("/api/v1")
	middlewares := middleware.GetMiddleware(serverInstance)

	// TODO: Figure out whether I should provide the server.Services or just one by one
	// Handlers
	userHandler := handler.GetUserHandler(serviceUser)
	authHandler := handler.GetAuthHandler(serviceSession, serviceAuth)

	// Routes
	userHandler.RegisterRoutes(middlewares, apiV1)
	authHandler.RegisterRoutes(middlewares, apiV1)

	serverInstance.StartDebug()
}
