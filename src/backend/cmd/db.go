package main

import (
	"fmt"

	database "../internal/database"
	gorm "github.com/jinzhu/gorm"
)

// dbConnect connects to the database
func dbConnect() (*database.DB, error) {
	psqlInfo := fmt.Sprintf("host=db port=5432 user=ggp-user password=password dbname=development_db sslmode=disable")
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	return &database.DB{DB: db}, nil
}
