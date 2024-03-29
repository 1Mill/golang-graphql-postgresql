package database

import (
	"context"
	"fmt"
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

// Connect to the database instance
func Connect() (*DB, error) {
	psqlInfo := fmt.Sprintf("host=db port=5432 user=ggp-user password=password dbname=development_db sslmode=disable")
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
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

// User fetches the user from the database by their id
func (db *DB) User(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
