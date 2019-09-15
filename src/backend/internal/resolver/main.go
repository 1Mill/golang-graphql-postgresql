package resolver

import (
	"context"

	database "../database"
)

// Resolver is the root resolver for the graphql query
type Resolver struct {
	DB *database.DB
}

// Hello resolves hello query
func (_ *Resolver) Hello() string {
	return "Hello, World!"
}

// Testing resolves testing query
func (_ *Resolver) Testing() string {
	return "Testing"
}

// GetUser resolves the graphql query
func (r *Resolver) GetUser(ctx context.Context) (*UserResolver, error) {
	user, err := r.DB.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	s := UserResolver{
		db: r.DB,
		m:  *user,
	}
	return &s, nil
}
