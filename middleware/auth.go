package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("Authorization") // tokenString and sessionID is same thing
		if sessionID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session."})
			return
		}

		claims, err := utils.ValidateAccessToken(sessionID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("userID", claims.UserID) // store the userID in the context so handlers can grab it
		c.Set("sessionID", claims.SessionID)
		c.Set("role", claims.Role)
		c.Next() // Let the req continue to the header
	}
}
