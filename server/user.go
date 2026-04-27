package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/handler"
	"github.com/prakhar0009/go-todo/middleware"
)

func SetupUserRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	// Public routes
	v1.POST("/register", handler.Register)
	v1.POST("/login", handler.Login)

	//Protected routes
	v1.Use(middleware.AuthMiddleware())
	{
		v1.PUT("/logout", handler.Logout)
		//v1.PUT("/user/role", handler.SwitchRole)
		// using this route a user can be admin
	}
}
