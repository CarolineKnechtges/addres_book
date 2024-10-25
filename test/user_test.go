// tests/user_test.go
package tests

import (
	"address_book/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"address_book/config"
	"address_book/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	InitTestEnvironment(t) // 初始化工作目录

	// 初始化数据库连接
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// 添加测试数据
	config.DB.Create(&models.User{Name: "John", Email: "john@example.com"})

	// 设置路由
	router := routes.SetupRouter()

	// 创建 HTTP 请求
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 验证响应
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John")

	// 清理测试数据
	config.DB.Exec("DELETE FROM users WHERE 1")
}
