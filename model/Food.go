package model

import "github.com/neelance/graphql-go"

type Food struct {
	ID graphql.ID `json:"id" bson:"_id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Image string `json:"image"`
	Season string `json:"season"`
}
