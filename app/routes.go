package app

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router
	router.GET("/greet", s.greet())

	// Snippet routes
	snippetGroup := router.Group("/snippet")
	{
		snippetGroup.POST("/create/:projectId", s.createSnippetRequestHandler())
		snippetGroup.POST("/delete/:snippetId", s.deleteSnippetRequestHandler())
		snippetGroup.GET(":snippetId", s.getSnippetRequestHandler())
	}

	return router
}
