package resolver

import (
	"context"
	"fmt"

	database "../database"
	model "../model"
	graphql "github.com/graph-gophers/graphql-go"
)

// BookResolver contains the database and the Book model to resolve the graphql query against
type BookResolver struct {
	db *database.DB
	m  model.Book
}

// Book finds the book by their id
func (r *Resolver) Book(ctx context.Context, args struct{ ID graphql.ID }) (*BookResolver, error) {
	book, err := r.DB.Book(ctx, fmt.Sprint(args.ID))
	if err != nil {
		return nil, err
	}

	s := BookResolver{
		db: r.DB,
		m:  *book,
	}
	return &s, nil
}

// Author resolves the author graphql relationship
func (r *BookResolver) Author(ctx context.Context) (*UserResolver, error) {
	u, err := r.db.User(ctx, fmt.Sprint(r.m.AuthorID))
	if err != nil {
		return nil, err
	}
	return &UserResolver{db: r.db, m: *u}, nil
}

// DatePublished resolves the datePublished graphql attribute
func (r *BookResolver) DatePublished(ctx context.Context) *string {
	s := r.m.DatePublished
	if s == "" {
		return nil
	}
	return &s
}

// ID resolves the ID graphql attribute
func (r *BookResolver) ID(ctx context.Context) graphql.ID {
	s := graphql.ID(fmt.Sprint(r.m.ID))
	return s
}

// Title resolves the title graphql attribute
func (r *BookResolver) Title(ctx context.Context) string {
	s := r.m.Title
	return s
}
