package resolver

import (
	"context"
	"fmt"

	database "../database"
	model "../model"
	graphql "github.com/graph-gophers/graphql-go"
)

// UserResolver contains the database and the user model to resolve the graphql query against
type UserResolver struct {
	db *database.DB
	m  model.User
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

// Email resolves the user email
func (u *UserResolver) Email(ctx context.Context) string {
	s := u.m.Email
	return s
}

// ID resolves the user ID
func (u *UserResolver) ID(ctx context.Context) graphql.ID {
	s := graphql.ID(fmt.Sprint(u.m.ID))
	return s
}

// Name resolves the full name of the user
func (u *UserResolver) Name(ctx context.Context) *string {
	s := fmt.Sprintf("%s %s", u.m.NameFirst, u.m.NameLast)
	if s == "" {
		return nil
	}
	return &s
}

// NameFirst resolves the first name of the user
func (u *UserResolver) NameFirst(ctx context.Context) *string {
	s := u.m.NameFirst
	if s == "" {
		return nil
	}
	return &s
}

// NameLast resolves the last name of the user
func (u *UserResolver) NameLast(ctx context.Context) *string {
	s := u.m.NameLast
	if s == "" {
		return nil
	}
	return &s
}
