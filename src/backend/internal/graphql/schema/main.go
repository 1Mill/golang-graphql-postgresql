package schema

import "io/ioutil"

// GetSchema exports the graphql schema for the file
func GetSchema() (string, error) {
	b, err := ioutil.ReadFile("./graphql/schema/schema.graphql")
	return string(b), err
}
