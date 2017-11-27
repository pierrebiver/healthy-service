package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	localGraphql "github.com/healthy-service/graphql"
	"log"
	"github.com/graphql-go/handler"
	"github.com/mnmtanish/go-graphiql"
	"github.com/rs/cors"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: localGraphql.QueryType,
})

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	http.Handle("/graphql", c.Handler(h))
	http.HandleFunc("/graphiql", graphiql.ServeGraphiQL)

	log.Println("Starting GraphQL Server oYeah on http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
