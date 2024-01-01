package server

import (
	"github.com/labstack/echo/v4"
	"log"
	"server/core/ports"
	"server/utils"
)

func GetServerInstance(services *Services) *Server {

	echoInstance := echo.New()
	echoInstance.IPExtractor = echo.ExtractIPFromXFFHeader()

	server := &Server{
		Echo:     echoInstance,
		Services: services,
	}

	server.InitializeValidator()
	return server
}

type Server struct {
	Echo     *echo.Echo
	Services *Services
}

// Services This isn't a part of the service packages as one might expect
// because this defines what ports WE want to interact with, not all the ports present in the application
// Though in this case it is synonymous
type Services struct {
	ServiceSession ports.ServiceSession
	ServiceUser    ports.ServiceUser
}

func (s *Server) StartDebug() {
	s.Echo.Logger.Info(s.Echo.Start(utils.GetEnv("API_ADDRESS")))
	//s.Echo.Logger.Info(s.Echo.Start(":1234"))

}

func (s *Server) Start() {
	log.Fatal("NOT IMPLEMENTED")
}
