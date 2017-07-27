package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/healthy-service/graphql/type"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"all": &graphql.Field{
			Type: _type.FoodType[],
			Resolve: func(p graphql.ResolveParams) (interface{}, error){
				
			}
		}
	}
})
