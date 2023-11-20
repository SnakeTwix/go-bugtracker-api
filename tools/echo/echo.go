package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoInstance *echo.Echo

func GetEchoInstance() *echo.Echo {
	if echoInstance != nil {
		return echoInstance
	}

	echoInstance = echo.New()
	echoInstance.Use(middleware.CORS())

	InitializeValidator(echoInstance)

	return echoInstance
}
