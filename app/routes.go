package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	router.GET("/greet", s.greet())

	return router
}
