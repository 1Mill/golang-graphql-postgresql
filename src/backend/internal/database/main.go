package database

import (
	"context"
	"log"

	model "../model"
	gorm "github.com/jinzhu/gorm"
)

var users = []model.User{
	{
		Email:     "testing@html.erb",
		NameFirst: "First",
		NameLast:  "Last",
	},
	{
		Email: "missing_names@html.erb",
	},
	{
		Email:     "john@html.erb",
		NameFirst: "John",
	},
}

// DB is the root database for all operations
type DB struct {
	DB *gorm.DB
}

// Seed the database with test data
func (db *DB) Seed() {
	db.DB.DropTableIfExists(&model.User{})
	db.DB.AutoMigrate(&model.User{})

	for _, u := range users {
		err := db.DB.Create(&u).Error
		if err != nil {
			log.Panic(err)
		}
	}
}

// GetUser fetches the user from the database
func (db *DB) GetUser(ctx context.Context) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
