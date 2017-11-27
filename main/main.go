package main

import (
	"net/http"
	"log"
	localGraph "github.com/healthy-service/graphql"
	"github.com/neelance/graphql-go"
	"github.com/rs/cors"
	"github.com/mnmtanish/go-graphiql"
	"github.com/neelance/graphql-go/relay"
)

var schema *graphql.Schema

func init() {
	schema = graphql.MustParseSchema(localGraph.Schema, &localGraph.Resolver{})
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	http.Handle("/graphql", c.Handler(&relay.Handler{Schema: schema}))
	http.HandleFunc("/graphiql", graphiql.ServeGraphiQL)

	log.Println("Starting GraphQL Server on http://localhost:8080/")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
