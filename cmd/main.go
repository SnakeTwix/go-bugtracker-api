package main

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"net/http"
	"server/internal/adapters/handler"
	"server/internal/adapters/repository"
	"server/internal/core/services"
	"time"
)

func main() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.CORS())

	db := initBun()

	repoUser := repository.GetRepoUser(db)
	serviceUser := services.GetServiceUser(repoUser)
	userHandler := handler.GetUserHandler(serviceUser)

	userHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1234"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func initBun() *bun.DB {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:5437"),
		pgdriver.WithUser("user"),
		pgdriver.WithPassword("pg-pass"),
		pgdriver.WithDatabase("dev"),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqlDB := sql.OpenDB(pgconn)

	db := bun.NewDB(sqlDB, pgdialect.New())

	return db
}
