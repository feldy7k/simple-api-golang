package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api/models"
	"go-api/services"
	"go-api/utils"
	"net/http"
)

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Produce json
// @Param auth body models.Auth true "Auth data"
// @Success 200
// @Router /auth [post]
func AuthLogin(c *gin.Context) {
    var input models.Auth

	users, errG := services.GetUsers()

	if errG != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch users"})
		return
	}

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    for _, u := range users {
        if u.Email == input.Email {
            token, _ := utils.GenerateJWT(u.Email)
            c.JSON(http.StatusOK, gin.H{"message": "Login success.", "token": token})
            return
        }
    }

    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}