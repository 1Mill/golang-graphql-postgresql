package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User model
type User struct {
	gorm.Model

	email     string
	firstName string
	lastName  string
}

func main() {
	// * Build connection string
	psqlInfo := fmt.Sprintf("host=db port=5432 user=ggp-user password=password dbname=development_db sslmode=disable")

	// * Try to connect to the database
	fmt.Println("connecting to database")
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// * Try pinging the database
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to database")

	// * Drop User table if table already exists (to reset the database to clean state)
	db.DropTableIfExists(&User{})

	// * Create User table
	db.CreateTable(&User{})

	// // * Migrate schema
	// db.AutoMigrate(&User{})

	// // * Create db
	// db.Create(&User{email: "abc@html.erb", firstName: "Example", lastName: "Testing"})

	// // * Read
	// var user User
	// db.First(&user, 1)                           // * Get User with id 1
	// db.First(&user, "email = ?", "abc@html.erb") // * Find User with email

	// // * Print users to screen
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", user.email)
	// })

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
