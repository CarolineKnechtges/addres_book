// routes/router.go
package routes

import (
	"address_book/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	usersPath  = "/users"
	userIDPath = "/users/:id"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 配置 CORS 中间件
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 定义路由
	router.GET(usersPath, controllers.GetUsers)
	router.GET(userIDPath, controllers.GetUser)
	router.POST(usersPath, controllers.CreateUser)
	router.PUT(userIDPath, controllers.UpdateUser)
	router.DELETE(userIDPath, controllers.DeleteUser)

	return router
}
