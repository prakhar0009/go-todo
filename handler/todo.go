package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/database/dbHelper"
)

func CreateTodo(c *gin.Context) {
	sessionID := c.GetHeader("Authorization")

	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, err := dbHelper.GetUserBySession(sessionID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoID, err := dbHelper.CreateTodo(userID, input.Title, input.Description, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "todoID created successfully",
		"todoID":  todoID,
	})
}
