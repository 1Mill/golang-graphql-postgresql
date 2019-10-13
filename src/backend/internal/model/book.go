package model

import (
	gorm "github.com/jinzhu/gorm"
)

// Book model defines books within the application
type Book struct {
	gorm.Model

	DatePublished string
	Title         string

	// belongs_to relationships
	AuthorID uint
}
