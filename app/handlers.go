package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func (s *Server) greet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "up and running",
			"data":   "greetings!",
		}

		c.JSON(http.StatusOK, response)
	}
}

// POST
func (s *Server) createSnippetRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		// Param checking and calls to service here
		fmt.Printf("ProjectId %s \n", c.Param("projectId"))

		c.JSON(http.StatusOK, make(map[string]string))
	}
}

// POST
func (s *Server) deleteSnippetRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		// Param checking and calls to service here
		fmt.Printf("SnippetId %s \n", c.Param("snippetId"))

		c.JSON(http.StatusOK, make(map[string]string))
	}
}

// GET
func (s *Server) getSnippetRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		// Param checking and calls to service here

		response := map[string]string{
			"status": "success",
			"data":   "greetings!",
		}

		c.JSON(http.StatusOK, response)
	}
}
