package dgraph

import (
	"context"
	"io/ioutil"
	"log"

	graphql "github.com/graph-gophers/graphql-go"
)

type QueryParams struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func LoadSchema(path string) *graphql.Schema {
	schema, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	root, err := NewRoot()
	if err != nil {
		log.Fatal(err)
	}

	return graphql.MustParseSchema(string(schema), root)
}

func ParseQuery(ctx context.Context, schema *graphql.Schema, params QueryParams) *QueryBuilder {
	b := NewBuilder()
	ctx = context.WithValue(ctx, "builder", b)
	result := schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	if len(result.Errors) > 0 {
		for _, err := range result.Errors {
			b.Error(err.Message)
		}
	}
	return b
}
