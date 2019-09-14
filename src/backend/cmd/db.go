package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// DB is the database for this application
type DB struct {
	DB *gorm.DB
}

// dbConnect connects to the database
func dbConnect() (*DB, error) {
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

var users = []User{
	{
		Email: "testing@html.erb",
	},
	{
		Email: "sdlkfjlskf@html.erb",
	},
	{
		Email: "email@html.erb",
	},
}

// Seed the database with test data
func (db *DB) Seed() {
	db.DB.DropTableIfExists(&User{})
	db.DB.AutoMigrate(&User{})

	for _, u := range users {
		err := db.DB.Create(&u).Error
		if err != nil {
			log.Panic(err)
		}
	}
}
