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

		var userStatus struct {
			IsSuspended bool `db:"is_suspended"`
		}

		// Query to see if the session is valid and get the user's current suspension status
		query := `
			SELECT u.is_suspended 
			FROM users u
			JOIN user_session us ON u.id = us.user_id
			WHERE us.id = $1 AND us.archived_at IS NULL`

		err = database.Todo.Get(&userStatus, query, claims.SessionID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session is invalid or logged out"})
			return
		}

		if userStatus.IsSuspended {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Your account is suspended. Access denied."})
			return
		}

		c.Set("userID", claims.UserID) // store the userID in the context so handlers can grab it
		c.Set("sessionID", claims.SessionID)
		c.Set("role", claims.Role)
		c.Next() // Let the req continue to the header
	}
}
