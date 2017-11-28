package graphql

import (
	"github.com/healthy-service/model"
	"github.com/healthy-service/mongo"
	"github.com/neelance/graphql-go"
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
		foodUpdate(food: FoodInput!): Food!
	}

	input FoodInput {
		id: ID!
		name: String!
		category: String!
		season: String!
	}

	type Food {
		id: ID!
		name: String!
		category: String!
		season: String!
	}
`

type Resolver struct{}

type FoodInput struct {
	ID       graphql.ID
	Name     string
	Category string
	Season   string
}

type foodResolver struct {
	Food model.Food
}

type foodArgs struct {
	Food *FoodInput
}

func (f *foodResolver) ID() graphql.ID {
	return f.Food.ID
}

func (f *foodResolver) Name() string {
	return f.Food.Name
}

func (f *foodResolver) Season() string {
	return f.Food.Season
}

func (f *foodResolver) Category() string {
	return f.Food.Category
}

func (r *Resolver) Foods() *[]*foodResolver {
	foods, err := mongo.GetAllFoods()
	if err != nil {
		log.Fatal(err)
	}
	foodsResolver := make([]*foodResolver, 0)
	for _, f := range foods {
		foodsResolver = append(foodsResolver, &foodResolver{f})
	}
	return &foodsResolver
}

func (r *Resolver) FoodUpdate(args foodArgs) *foodResolver {
	input := args.Food
	newFood, err := mongo.UpdateFood(model.Food{ID: input.ID, Name: input.Name, Category: input.Category, Season: input.Season})
	if err != nil {
		log.Fatal("Cannot update food", err)
	}

	return &foodResolver{newFood}
}
