package main

import (
	"context"

	"github.com/jinzhu/gorm"
)

// GetUser fetches the user from the database
func (db *DB) GetUser(ctx context.Context) (string, error) {
	return "TESTING", nil
}
