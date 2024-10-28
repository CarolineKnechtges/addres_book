// Package repository/user_repository.go
package repository

import (
	"address_book/config"
	"address_book/models"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(id string) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

// CreateUser inserts a new user into the database
func CreateUser(user models.User) error {
	result := config.DB.Create(&user)
	return result.Error
}

// UpdateUserByID updates an existing user in the database by ID
func UpdateUserByID(id string, user models.User) error {
	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return err
	}
	// 更新用户字段
	existingUser.Name = user.Name
	existingUser.Email = user.Email // 假设你有这些字段
	// 继续更新其他字段...

	return config.DB.Save(&existingUser).Error
}

// DeleteUserByID removes a user from the database by ID
func DeleteUserByID(id string) error {
	result := config.DB.Delete(&models.User{}, id)
	return result.Error
}
