package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	localGraphql "github.com/healthy-service/graphql"
	"encoding/json"
	"log"
	"fmt"
)

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		fmt.Println(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v \n", result.Errors)
	}
	return result
}

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    localGraphql.QueryType,
		Mutation: localGraphql.Mutation,
	})

	http.Handle("/graphql", handler(schema))

	log.Println("Starting GraphQL Server on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
