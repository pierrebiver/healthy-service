package graphql

import (
	"github.com/healthy-service/model"
	"github.com/healthy-service/mongo"
	"log"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {
		foods(): [Food]
	}

	type Mutation {
		foodUpdate(food: Food!): Food!
	}

	interface Food {
		id: ID!
		name: String!
		category: String!
		season: String!
		image: String
	}
`

type Resolver struct{}

type FoodResolver struct {
	Food *model.Food
}

// TODO func to access data of struc resolver

func (r *Resolver) Foods() *[]*FoodResolver {
	foods, err := mongo.GetAllFoods()
	if err != nil {
		log.Fatal(err)
	}
	foodsResolver := make([]*FoodResolver,0)
	for _,f := range foods {
		foodsResolver = append(foodsResolver,&FoodResolver{&f})
	}
	return &foodsResolver
}

func (r *Resolver) FoodUpdate(args *struct {
	Food model.Food
}) *FoodResolver {
	newFood, err := mongo.UpdateFood(args.Food)
	if err != nil {
		log.Fatal("Cannot update food", err)
	}

	return &FoodResolver{&newFood}
}
