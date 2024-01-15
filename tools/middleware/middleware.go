package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"server/core/domain"
	"server/core/enums/cookies"
	"server/tools/server"
	"server/utils"
)

type Middleware struct {
	server *server.Server
}

var middlewareInstance *Middleware

// GetMiddleware Gets the middleware struct and register the global middleware on the server
func GetMiddleware(server *server.Server) *Middleware {
	if middlewareInstance != nil {
		return middlewareInstance
	}

	middlewareInstance = &Middleware{
		server: server,
	}

	middlewareInstance.registerGlobalMiddleware()
	return middlewareInstance
}

func (m *Middleware) registerGlobalMiddleware() {
	m.server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:                             []string{utils.GetEnv("CLIENT_ADDRESS")},
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: false,
	}))

	middlewareInstance.server.Echo.Use(middlewareInstance.SetUserIfSession)
	middlewareInstance.server.Echo.Use(middlewareInstance.LogRequest)
}

func (m *Middleware) SetUserIfSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie(cookies.Session)
		if err != nil {
			//log.Print("No session cookie\n")
			return next(ctx)
		}

		session, err := m.server.Services.ServiceSession.GetSession(ctx.Request().Context(), cookie.Value)
		if err != nil {
			return next(ctx)
		}

		ctx.Set("session", session)
		return next(ctx)
	}
}

func (m *Middleware) CheckLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, ok := ctx.Get("session").(*domain.Session)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return next(ctx)
	}
}

func (m *Middleware) LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.Info(ctx.Request().RequestURI)

		return next(ctx)
	}
}
