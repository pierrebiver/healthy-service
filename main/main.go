package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	localGraphql "github.com/healthy-service/graphql"
	"log"
	"github.com/graphql-go/handler"
	"github.com/mnmtanish/go-graphiql"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    localGraphql.QueryType,
})

type GraphQLPost struct {
	query string
}

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", h)
	http.HandleFunc("/graphiql", graphiql.ServeGraphiQL)

	log.Println("Starting GraphQL Server on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
