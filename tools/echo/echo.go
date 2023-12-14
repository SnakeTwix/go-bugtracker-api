package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"server/utils"
)

var echoInstance *echo.Echo

func GetEchoInstance() *echo.Echo {
	if echoInstance != nil {
		return echoInstance
	}

	echoInstance = echo.New()
	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:                             []string{utils.GetEnv("CLIENT_ADDRESS")},
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: false,
	}))

	InitializeValidator(echoInstance)

	return echoInstance
}
