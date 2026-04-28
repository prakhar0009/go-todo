package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/handler"
	"github.com/prakhar0009/go-todo/middleware"
)

func SetupTodoRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		// User routes
		v1.POST("/todo", handler.CreateTodo)
		v1.GET("/todos", handler.GetTodos)
		v1.GET("/todo/:id", handler.GetTodoByID)
		v1.PUT("/todo/:id", handler.UpdateTodo)
		v1.DELETE("/todo/:id", middleware.AdminMiddleware(), handler.DeleteTodo)


		// Admin routes
		admin := v1.Group("/admin")
		admin.Use(middleware.AdminMiddleware()) // only admin can access
		{
			admin.GET("/todos", handler.GetTodos)
			admin.DELETE("/")
		}
	}
}
