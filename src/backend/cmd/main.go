package main

import (
	"fmt"
	"log"
	"net/http"

	"../pkg/graphql/schema"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=db port=5432 user=ggp-user password=password dbname=development_db sslmode=disable")
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Panic(err)
	}

	s := schema.String("./gql")
	graphqlSchema := graphql.MustParseSchema(s, &Query{})

	http.Handle("/graphql", &relay.Handler{Schema: graphqlSchema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
