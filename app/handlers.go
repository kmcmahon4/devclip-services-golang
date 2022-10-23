package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) greet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "greetings!",
		}

		c.JSON(http.StatusOK, response)
	}
}
