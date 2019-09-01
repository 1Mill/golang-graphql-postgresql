package schema

// Schema returns the graphql shema
func Schema() string {
	return `
		schema {
			query: Query
		}
		type Query {
			hello: String!
		}
	`
}
