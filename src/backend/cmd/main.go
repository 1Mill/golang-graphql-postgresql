package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	s "../internal/graphql/schema"
	graphql "github.com/graph-gophers/graphql-go"
)

type query struct{}

func main() {
	psqlInfo := fmt.Sprintf("host=db port=5432 user=ggp-user password=password dbname=development_db sslmode=disable")
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s.Schema(), &query{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (_ *query) Hello() string {
	return "Hello, World!"
}
