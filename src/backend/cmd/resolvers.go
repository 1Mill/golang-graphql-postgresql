package main

// Query placeholder
type Query struct{}

// Hello resolves hello query
func (_ *Query) Hello() string {
	return "Hello, World!"
}

// Testing resolves testing query
func (_ *Query) Testing() string {
	return "Testing"
}
