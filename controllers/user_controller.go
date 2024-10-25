// Package controllers/user_controller.go
package controllers

import (
	"address_book/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users, err := services.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
