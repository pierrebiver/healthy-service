package _type

import "github.com/graphql-go/graphql"

var RecipeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Recipe",
	Fields: graphql.Fields{
		"id":                      &graphql.Field{Type: graphql.ID},
		"name":                    &graphql.Field{Type: graphql.String},
		"description":             &graphql.Field{Type: graphql.String},
		"ingredients":             &graphql.Field{Type: graphql.NewList(FoodType)},
		"preparationTimeInMinute": &graphql.Field{Type: graphql.Int},
	},
})
