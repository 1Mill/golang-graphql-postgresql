package main

import "github.com/jinzhu/gorm"

// Resolver is the root resolver for the graphql query
type Resolver struct {
	db *gorm.DB
}

// Hello resolves hello query
func (_ *Resolver) Hello() string {
	return "Hello, World!"
}

// Testing resolves testing query
func (_ *Resolver) Testing() string {
	return "Testing"
}
