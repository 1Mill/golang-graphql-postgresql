package main

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// User model defins users within the application
type User struct {
	gorm.Model

	Email     string
	firstName string
	lastName  string
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

// UserResolver contains the database and the user model to resolve the graphql query against
type UserResolver struct {
	db *DB
	m  User
}

// ID resolves the user ID
func (u *UserResolver) ID(ctx context.Context) *graphql.ID {
	s := graphql.ID(fmt.Sprint(u.m.ID))
	return &s
}
