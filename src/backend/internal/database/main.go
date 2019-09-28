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
var books = []model.Book{
	{
		DatePublished: "1234-12-12",
		Title:         "Some random title",
	},
	{
		DatePublished: "",
		Title:         "Some crazy title for a book that is long!",
	},
	{
		DatePublished: "1929-02-01",
		Title:         "An old title",
	},
	{
		DatePublished: "1992-01-01",
		Title:         "A new title",
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
	db.DB.DropTableIfExists(&model.User{}, &model.Book{})
	db.DB.AutoMigrate(&model.User{}, &model.Book{})

	for _, u := range users {
		err := db.DB.Create(&u).Error
		if err != nil {
			log.Panic(err)
		}
	}
	for _, b := range books {
		err := db.DB.Create(&b).Error
		if err != nil {
			log.Panic(err)
		}
	}
}

// Book fetches a book from the database by their id
func (db *DB) Book(ctx context.Context, id string) (*model.Book, error) {
	var book model.Book
	err := db.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
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
