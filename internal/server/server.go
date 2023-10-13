package server

import (
	"server/internal/controllers"
	"server/internal/models"
)

type Server struct {
	ModelUsers     models.IModelUser
	ControllerUser controllers.IControllerUser
}

var server *Server

func GetServer() *Server {
	if server != nil {
		return server
	}

	server = &Server{
		ModelUsers:     models.ModelUser{},
		ControllerUser: controllers.ControllerUser{},
	}

	return server
}
