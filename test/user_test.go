// tests/user_test.go
package tests

import (
	"address_book/config"
	"address_book/models"
	"address_book/routes"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitDB() {
	// 设置工作目录为项目路径
	if err := os.Chdir("/Users/paoyou/GolandProjects/address_book"); err != nil {
		panic(fmt.Sprintf("Failed to change directory: %v", err))
	}

	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}
}

func CleanUpDB() {
	config.DB.Exec("DELETE FROM users WHERE 1")
}

func CreateUserRequest(user models.User) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	router := routes.SetupRouter()
	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)
	return recorder
}

func UpdateUserRequest(userID uint, user models.User) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	router := routes.SetupRouter()
	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", userID), bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)
	return recorder
}

func DeleteUserRequest(userID uint) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	router := routes.SetupRouter()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", userID), nil)
	router.ServeHTTP(recorder, req)
	return recorder
}

func TestUpdateUser(t *testing.T) {
	// 初始化数据库
	InitDB()
	defer CleanUpDB() // 在测试结束后清理数据库

	// 添加测试数据
	user := models.User{Name: "Bob", Email: "bob@example.com"}
	config.DB.Create(&user)

	updatedUser := models.User{Name: "Bob Updated", Email: "bob_updated@example.com"}
	recorder := UpdateUserRequest(user.ID, updatedUser)

	// 验证响应
	assert.Equal(t, http.StatusOK, recorder.Code)

	// 验证用户是否被更新
	var fetchedUser models.User
	config.DB.First(&fetchedUser, user.ID)
	assert.Equal(t, "Bob Updated", fetchedUser.Name)

}

func TestCreateUser(t *testing.T) {
	// 初始化数据库
	InitDB()
	defer CleanUpDB() // 在测试结束后清理数据库

	user := models.User{Name: "Alice", Email: "alice@example.com"}
	recorder := CreateUserRequest(user)

	// 验证响应
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// 验证用户是否在数据库中
	var createdUser models.User
	config.DB.Where("email = ?", "alice@example.com").First(&createdUser)
	assert.Equal(t, "Alice", createdUser.Name)

}

func TestGetUsers(t *testing.T) {
	// 初始化数据库
	InitDB()
	defer CleanUpDB() // 在测试结束后清理数据库

	recorder := httptest.NewRecorder()
	router := routes.SetupRouter()

	// 添加测试数据
	config.DB.Create(&models.User{Name: "John", Email: "john@example.com"})

	// 创建 HTTP 请求
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(recorder, req)

	// 验证响应
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "John")

}

func TestDeleteUser(t *testing.T) {
	// 初始化数据库
	InitDB()
	defer CleanUpDB() // 在测试结束后清理数据库

	// 添加测试数据
	user := models.User{Name: "Charlie", Email: "charlie@example.com"}
	config.DB.Create(&user)

	recorder := DeleteUserRequest(user.ID)

	// 验证响应
	assert.Equal(t, http.StatusNoContent, recorder.Code)

	// 验证用户是否已被删除
	var deletedUser models.User
	result := config.DB.Where("id = ?", user.ID).First(&deletedUser)
	assert.True(t, result.Error != nil) // 确保用户已被删除
}
