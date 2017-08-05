package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/healthy-service/graphql/type"
	"github.com/healthy-service/mongo"
	"fmt"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"foods": &graphql.Field{
			Type:        graphql.NewList(_type.FoodType),
			Description: "Gel all foods",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				foods, err := mongo.GetAllFoods()

				return foods, err
			},
		},
		"recipes": &graphql.Field{
			Type:        graphql.NewList(_type.RecipeType),
			Description: "get All recipes",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				recipes, err := mongo.GetAllRecipes()
				if err != nil {
					fmt.Print(err)
				}
				return recipes, err
			},
		},
	},
})
