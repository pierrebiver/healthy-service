package _type

import "github.com/graphql-go/graphql"

func Get()  {
	 graphql.NewObject(graphql.ObjectConfig{
		Name: "Food",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.String},
			"name":     &graphql.Field{Type: graphql.String},
			"image":    &graphql.Field{Type: graphql.String},
			"category": &graphql.Field{Type: graphql.String},
			"season":   &graphql.Field{Type: graphql.List{}},
		},
	})
}
