package main

import (
	"fmt"

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
