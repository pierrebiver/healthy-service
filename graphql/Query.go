package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/healthy-service/graphql/type"
	"github.com/healthy-service/mongo"
	"github.com/healthy-service/model"
	"gopkg.in/mgo.v2/bson"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"all": &graphql.Field{
			Type:        graphql.NewList(_type.FoodType),
			Description: "Gel all foods",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				session, _ := mongo.GetSession()
				defer session.Close()

				db := session.DB(mongo.DBName)

				foods := make([]model.Food, 0)
				err := db.C("foods").Find(bson.M{}).All(&foods)
				return foods, err
			},
		},
	},
})
