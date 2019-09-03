package main

// Query placeholder
type Query struct{}

func (_ *Query) Hello() string {
	return "Hello, World!"
}

func (_ *Query) Testing() string {
	return "SLKFJLKSDJFLJDSLKFSF"
}
