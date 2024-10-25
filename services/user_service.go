// Package services/user_service.go
package services

import (
	"address_book/models"
	"address_book/repository"
)

func FetchAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}
