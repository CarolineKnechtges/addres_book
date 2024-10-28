// Package services/user_service.go
package services

import (
	"address_book/models"
	"address_book/repository"
)

// FetchAllUsers retrieves all users from the repository
func FetchAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

// FetchUserById retrieves a user by ID from the repository
func FetchUserById(id string) (models.User, error) {
	return repository.GetUserByID(id)
}

// AddUser adds a new user to the repository
func AddUser(user models.User) error {
	return repository.CreateUser(user)
}

// UpdateUser updates an existing user in the repository
func UpdateUser(id string, user models.User) error {
	return repository.UpdateUserByID(id, user)
}

// DeleteUser removes a user from the repository by ID
func DeleteUser(id string) error {
	return repository.DeleteUserByID(id)
}
