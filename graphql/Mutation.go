package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/healthy-service/graphql/type"
	"github.com/healthy-service/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/healthy-service/mongo"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"foodCreate": &graphql.Field{
			Type:        _type.FoodType,
			Description: "Add new food",
			Args: graphql.FieldConfigArgument{
				"food": &graphql.ArgumentConfig{
					Type: _type.FoodType,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				food, _ := params.Args["food"].(model.Food)

				food.ID = bson.NewObjectId().String()
				newFood, err := mongo.AddFood(food)

				return newFood, err
			},
		},
		"foodUpdate": &graphql.Field{
			Type:        _type.FoodType,
			Description: "update food",
			Args: graphql.FieldConfigArgument{
				"food": &graphql.ArgumentConfig{
					Type: _type.FoodType,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				food, _ := params.Args["food"].(model.Food)

				newFood, err := mongo.UpdateFood(food)

				return newFood, err
			},
		},
	},
})
