package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization") // tokenString and tokenString is same thing
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required."})
			return
		}

		claims, err := utils.ValidateAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		var exists bool
		query := `SELECT EXISTS(SELECT 1 FROM user_session WHERE id = $1 AND archived_at IS NULL)`
		err = database.Todo.Get(&exists, query, claims.SessionID)

		if err != nil || !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session has been logged out"})
			return
		}
		c.Set("userID", claims.UserID) // store the userID in the context so handlers can grab it
		c.Set("sessionID", claims.SessionID)
		c.Set("role", claims.Role)
		c.Next() // Let the req continue to the header
	}
}
