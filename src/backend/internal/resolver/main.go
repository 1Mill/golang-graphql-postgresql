package resolver

import (
	database "../database"
)

// Resolver is the root resolver for the graphql query
type Resolver struct {
	DB *database.DB
}
