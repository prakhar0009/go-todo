package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/handler"
)

func SetupUserRoutes(r *gin.Engine) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	//authorized := r.Group("/")
	//authorized.Use(middleware.AuthRequired()){
	//	authorized.POST("/logout", handler.Logout)
	//}
}
