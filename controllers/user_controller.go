package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api/models"
	"go-api/services"
	"net/http"
	"github.com/google/uuid"
)

// GetUsers godoc
// @Summary Get users list
// @Description get users list
// @Tags users
// @Produce json
// @Success 200
// @Router /users [get]
func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
// GetUser godoc
// @Summary Get users by ID
// @Description get users by ID
// @Tags users
// @Produce json
// @Success 200
// @Router /users/:id [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary      Create user
// @Description  Create new user
// @Tags         users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uuid.New().String()
	err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CreateUser failed"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user
// @Tags         users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200
// @Router       /users [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UpdateUser failed"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// GetUser godoc
// @Summary      Delete user
// @Description  Delete user
// @Tags         users
// @Produce      json
// @Success 200
// @Router       /users/:id [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := services.RemoveUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DeleteUser failed"})
		return
	}
	c.JSON(http.StatusOK, id)
}