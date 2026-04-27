package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prakhar0009/go-todo/database/dbHelper"
	"github.com/prakhar0009/go-todo/models"
	"github.com/prakhar0009/go-todo/utils"
)

func Register(c *gin.Context) {

	var input models.CreateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	isUserExist, err := dbHelper.IsUserExist(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if isUserExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
		return
	}

	hashedPass, _ := utils.HashPassword(input.Password)
	err = dbHelper.CreateUser(input.Email, input.Username, hashedPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {

	var input models.LoginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Printf("Binding Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})
		return
	}
	fmt.Printf("Attempting login for: %s with password: %s\n", input.Email, input.Password)

	user, err := dbHelper.GetUserByEmail(input.Email)
	if err != nil {
		// This will print the actual database error to your terminal
		fmt.Printf("Database Error for %s: %v\n", input.Email, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		fmt.Println("Password verification failed for user:", input.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	expiry := time.Now().Add(60 * 24 * time.Hour)
	sessionID, err := dbHelper.CreateUserSession(user.ID, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create login session",
		})
		return
	}

	//c.SetCookie("session_token", sessionID, 60*24*3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message":    "User logged in successfully",
		"session_id": sessionID,
	})
}

func Logout(c *gin.Context) {
	sessionID := c.GetHeader("Authorization")
	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No active session"})
		return
	}
	err := dbHelper.DeleteUserSession(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user session",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}

func SwitchRole(c *gin.Context) {
	userID := c.GetString("userID")

	var input struct {
		Role models.UserRole `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, Role required"})
	}
	if input.Role != models.RoleUser && input.Role != models.RoleAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role must be one of admin or user"})
	}
	if err := dbHelper.UpdateUserRole(userID, input.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Role switched successfully",
		"role":    input.Role,
	})
}
