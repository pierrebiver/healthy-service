package model



type Recipe struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Ingredients []string `json:"ingredients"`
	PreparationTimeInMinute int `json:"preparationTimeInMinute"`
}
