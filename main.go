package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"hello": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "world", nil
					},
				},
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to create new schema: %s", err)
	}

	query := `
		{
			hello
		}
	`
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if r.HasErrors() {
		log.Fatalf("failed to execute graphql operation: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s\n", rJSON)
}
