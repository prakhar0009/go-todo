package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/database/dbHelper"
	"github.com/prakhar0009/go-todo/models"
)

func CreateTodo(c *gin.Context) {
	userID := c.GetString("userID")

	var input models.CreateTodo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Title and Description are required."})
		return
	}

	layout := "2006-01-02 15:04:05"

	expiry, err := time.Parse(layout, input.ExpiresAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format."})
	}

	todoID, err := dbHelper.CreateTodo(userID, input.Title, input.Description, expiry)
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
	userID := c.GetString("userID")

	statusFilter := c.Query("status")

	todos, err := dbHelper.GetAllTodo(userID, statusFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks."})
		return
	}

	for i := range todos {
		todos[i].SyncStatus()
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {

	userID := c.GetString("userID")

	todoID := c.Param("id")

	todo, err := dbHelper.GetTodoBYID(userID, todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or Unauthorized user"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {

	userID := c.GetString("userID")

	todoID := c.Param("id")

	var input models.UpdateTodo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	existingTodo, err := dbHelper.GetTodoBYID(userID, todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	if input.Title != "" {
		existingTodo.Title = input.Title
	}
	if input.Description != "" {
		existingTodo.Description = input.Description
	}
	if input.IsCompleted != nil {
		existingTodo.IsCompleted = *input.IsCompleted
	}

	existingTodo.SyncStatus()

	updatedTodo, err := dbHelper.UpdateTodo(existingTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {

	userID := c.GetString("userID")

	todoID := c.Param("id")

	if err := dbHelper.DeleteTodo(userID, todoID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delete failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo archived successfully."})
}
