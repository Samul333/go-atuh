package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unqiue" json:"email"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`
}
