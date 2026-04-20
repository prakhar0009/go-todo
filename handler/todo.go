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
		// expiry_at
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

func GetAllTodo(c *gin.Context) {
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

	todos, err := dbHelper.GetAllTodo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
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

	todoID := c.Param("id")

	todo, err := dbHelper.GetTodoBYID(userID, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
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

	todoID := c.Param("id")
	var input struct {
		ID          string `json:"id" binding:"required"`
		Title       string `json:"title"`
		Description string `json:"description"`
		IsCompleted string `json:"is_completed"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	oldTodo, err := dbHelper.GetTodoBYID(userID, todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	if input.Title != "" {
		oldTodo.Title = input.Title
	}

	if input.Description != "" {
		oldTodo.Description = input.Description
	}

	if input.IsCompleted != "" {
		if input.IsCompleted == "true" {
			oldTodo.IsCompleted = true
		} else if input.IsCompleted == "false" {
			oldTodo.IsCompleted = false
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "is_completed must be 'true' or 'false' string only"})
			return
		}
	}

	updatedTodo, err := dbHelper.UpdateTodo(oldTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {
	sessionID := c.GetHeader("Authorization")
	userID, err := dbHelper.GetUserBySession(sessionID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	todoID := c.Param("id")

	err = dbHelper.DeleteTodo(userID, todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo archived successfully"})
}
