package server

import (
	"net/http"
	"storage_api/internal/config"

	"github.com/gin-gonic/gin"
)

// TODO: Handle error response with JSON and application/json header
func AuthMiddleware(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "missing authorization header")
			return
		}

		if authHeader != config.AuthToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "incorrect authorization header")
			return
		}

		c.Next()
	}
}
