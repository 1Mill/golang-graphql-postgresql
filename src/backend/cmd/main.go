package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DatabaseInfo is used to connect to the database
type DatabaseInfo struct {
	host     string
	name     string
	password string
	port     int
	user     string
}

// User model
type User struct {
	gorm.Model

	email     string
	firstName string
	lastName  string
}

func connectToDb() *gorm.DB {
	// * Database info object
	var databaseInfo DatabaseInfo
	databaseInfo.host = "db"
	databaseInfo.name = "development_db"
	databaseInfo.password = "password"
	databaseInfo.port = 5432
	databaseInfo.user = "ggp-user"

	// * Build connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseInfo.host, databaseInfo.port, databaseInfo.user, databaseInfo.password, databaseInfo.name)

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
	return db
}

func main() {
	db := connectToDb()

	// * Migrate schema
	db.AutoMigrate(&User{})

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
