// tests/user_test.go
package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"address_book/config"
	"address_book/models"
	"address_book/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetUsers(t *testing.T) {
	// 初始化数据库连接
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	// 添加测试数据
	config.DB.Create(&models.User{Name: "John", Email: "john@example.com"})

	// 设置路由
	router := SetUpRouter()
	router.GET("/users", routes.SetupRouter().Routes()[0].HandlerFunc) // 绑定路由

	// 创建 HTTP 请求
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 验证响应
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John")
}
