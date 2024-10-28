// Package routes/router.go
package routes

import (
	"address_book/controllers"
	"github.com/gin-gonic/gin"
)

const (
	usersPath  = "/users"
	userIDPath = "/users/:id"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 定义路由
	router.GET(usersPath, controllers.GetUsers)
	router.GET(userIDPath, controllers.GetUser)
	router.POST(usersPath, controllers.CreateUser)
	router.PUT(userIDPath, controllers.UpdateUser)
	router.DELETE(userIDPath, controllers.DeleteUser)

	return router
}
