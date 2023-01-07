package app

import (
	devclip_service "app/pkg/devclip"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine, devclipService devclip_service.DevclipService) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run() error {
	// run the server through the router
	return s.Routes().Run(":8080")
}
