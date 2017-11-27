package model

type Food struct {
	ID string `json:"id" bson:"_id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Image string `json:"image"`
	Season string `json:"season"`
}
