package main

import (
	"log"
	"net/http"

	"../pkg/graphql/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is the root database for all operations
func main() {
	db, err := dbConnect()
	if err != nil {
		log.Panic(err)
	}

	db.Seed()

	s := schema.String("./gql")
	graphqlSchema := graphql.MustParseSchema(s, &Resolver{db: db})

	http.Handle("/graphql", &relay.Handler{Schema: graphqlSchema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
