package main

import (
	"log"
	"net/http"

	resolver "../internal/resolver"
	schema "../pkg/graphql/schema"
	graphql "github.com/graph-gophers/graphql-go"
	relay "github.com/graph-gophers/graphql-go/relay"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := dbConnect()
	if err != nil {
		log.Panic(err)
	}

	db.Seed()

	s := schema.String("./gql")
	graphqlSchema := graphql.MustParseSchema(s, &resolver.Resolver{DB: db})

	http.Handle("/graphql", &relay.Handler{Schema: graphqlSchema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
