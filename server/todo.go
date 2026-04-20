package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/handler"
)

func SetupTodoRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/todo", handler.CreateTodo)
		v1.GET("/todos", handler.GetAllTodo)
		v1.GET("/todo/:id", handler.GetTodoByID)
		v1.PUT("/todo/:id", handler.UpdateTodo)
		v1.DELETE("/todo/:id", handler.DeleteTodo)
	}
}
