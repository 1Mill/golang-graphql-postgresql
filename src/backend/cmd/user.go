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
	NameFirst string
	NameLast  string
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

// Email resolves the user email
func (u *UserResolver) Email(ctx context.Context) string {
	s := u.m.Email
	return s
}

// ID resolves the user ID
func (u *UserResolver) ID(ctx context.Context) graphql.ID {
	s := graphql.ID(fmt.Sprint(u.m.ID))
	return s
}

// NameFirst resolves the first name of the user
func (u *UserResolver) NameFirst(ctx context.Context) *string {
	s := u.m.NameFirst
	if s == "" {
		return nil
	}
	return &s
}

// NameLast resolves the last name of the user
func (u *UserResolver) NameLast(ctx context.Context) *string {
	s := u.m.NameLast
	if s == "" {
		return nil
	}
	return &s
}
