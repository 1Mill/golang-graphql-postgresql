package schema

import "io/ioutil"

// String exports the graphql schema as a string to be consumed by graphql-go
func String() (string, error) {
	b, err := ioutil.ReadFile("./graphql/schema/schema.graphql")
	return string(b), err
}
