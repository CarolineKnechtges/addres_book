// Package models/user.go
package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	TEL   string `json:"tel"`
}
