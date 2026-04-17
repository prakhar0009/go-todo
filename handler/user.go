package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/database/dbHelper"
	"github.com/prakhar0009/go-todo/utils"
)

func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExist, err := dbHelper.IsUserExist(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
		return
	}

	hashedPass, _ := utils.HashPassword(input.Password)
	err = dbHelper.CreateUser(input.Email, input.Username, hashedPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := dbHelper.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	expiry := time.Now().Add(60 * 24 * time.Hour)
	sessionID, err := dbHelper.CreateUserSession(user.ID, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//c.SetCookie("session_token", sessionID, 60*24*3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User logged in", "session_id": sessionID})
}

func Logout(c *gin.Context) {
	sessionID := c.GetHeader("Authorization")
	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}
	err := dbHelper.DeleteUserSession(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged out", "session_id": sessionID})
}
