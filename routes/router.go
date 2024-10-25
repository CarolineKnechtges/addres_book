// Package routes/router.go
package routes

import (
	"address_book/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 定义路由
	router.GET("/users", controllers.GetUsers)

	return router
}
