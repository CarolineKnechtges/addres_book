// main.go
package main

import (
	"address_book/config"
	"address_book/models"
	"address_book/routes"
)

func main() {
	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移数据库
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
		return
	}

	// 设置路由
	router := routes.SetupRouter()

	// 启动服务器
	err = router.Run(":8888")
	if err != nil {
		panic(err)
		return
	}
}
