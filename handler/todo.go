package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/database/dbHelper"
	"github.com/prakhar0009/go-todo/models"
)

func CreateTodo(c *gin.Context) {
	sessionID := c.GetHeader("Authorization")

	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, err := dbHelper.GetUserBySession(sessionID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}
	//var input struct {
	//	Title       string `json:"title" binding:"required"`
	//	Description string `json:"description" binding:"required"`
	//	// expiry_at
	//}

	var input models.CreateTodo

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Title and Description are required."})
		return
	}

	todoID, err := dbHelper.CreateTodo(userID, input.Title, input.Description /*input.ExpiresAt*/, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	todos, err := dbHelper.GetAllTodo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get all todos"})
		return
	}
	for i := range todos {
		todos[i].SyncStatus()
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user session"})
		return
	}

	todoID := c.Param("id")

	todo, err := dbHelper.GetTodoBYID(userID, todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or Unauthorized user"})
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
	//var input struct {
	//	ID          string `json:"id" binding:"required"`
	//	Title       string `json:"title"`
	//	Description string `json:"description"`
	//	IsCompleted string `json:"is_completed"`
	//}

	var input models.UpdateTodo
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format provided"})
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

	if input.IsCompleted != nil {
		oldTodo.IsCompleted = *input.IsCompleted
	}

	oldTodo.SyncStatus()

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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access."})
		return
	}

	todoID := c.Param("id")

	err = dbHelper.DeleteTodo(userID, todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find the Todo to delete."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo archived successfully."})
}
