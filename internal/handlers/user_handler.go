package handlers

import (
	"go-backend-app/internal/models"
	"go-backend-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
    users, err := services.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
	}

	err := services.AddUser(user)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func LoginUser(c *gin.Context) {
	var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
	}

	user, valid, err := services.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error logging in"})
			return
	}

	if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
	}

	c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"user":    user,
	})
}