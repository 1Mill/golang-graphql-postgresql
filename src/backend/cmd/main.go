package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	database "../internal/database"
	graphql "github.com/graph-gophers/graphql-go"
	relay "github.com/graph-gophers/graphql-go/relay"
	resolver "../internal/resolver"
	schema "../pkg/graphql/schema"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Panic(err)
	}

	db.Seed()

	s := schema.String("./gql")
	graphqlSchema := graphql.MustParseSchema(s, &resolver.Resolver{DB: db})

	http.Handle("/graphql", &relay.Handler{Schema: graphqlSchema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
