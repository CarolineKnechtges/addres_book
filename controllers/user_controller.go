// Package controllers/user_controller.go
package controllers

import (
	"address_book/models" // 假设你有一个模型文件定义用户结构体
	"address_book/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUsers returns all users
func GetUsers(c *gin.Context) {
	users, err := services.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser returns a single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id") // 获取用户ID
	user, err := services.FetchUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user models.User // 假设你的用户模型在models包中
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.AddUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id") // 获取用户ID
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateUser(id, user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id") // 获取用户ID

	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil) // 返回204 No Content
}
