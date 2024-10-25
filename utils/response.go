// Package utils/response.go
package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSONSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func JSONError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": message})
}
