// repository/user_repository.go
package repository

import (
	"address_book/config"
	"address_book/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}
