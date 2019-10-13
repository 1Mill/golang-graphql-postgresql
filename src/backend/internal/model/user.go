package model

import (
	gorm "github.com/jinzhu/gorm"
)

// User model defins users within the application
type User struct {
	gorm.Model

	Email     string
	NameFirst string
	NameLast  string

	// has_many relatinship
	Books []Book `gorm:"foreignkey:AuthorID"`
}
