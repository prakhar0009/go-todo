package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/handler"
)

func SetupTodoRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/create", handler.CreateTodo)
		//v1.GET()
		//v1.PUT()
		//v1.DELETE()
	}
}
