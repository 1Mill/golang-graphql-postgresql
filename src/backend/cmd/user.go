package main

import (
	"context"

	"github.com/jinzhu/gorm"
)

// User model defins users within the application
type User struct {
	gorm.Model

	Email     string
	firstName string
	lastName  string
}

// UserResolver contains the database and the user model to resolve the graphql query against
type UserResolver struct {
	db *DB
	m  User
}

// GetUser fetches the user from the database
func (db *DB) GetUser(ctx context.Context) (*User, error) {
	var user User
	err := db.DB.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
