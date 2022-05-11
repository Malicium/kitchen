package recipes

type Recipe struct {
	ID           string       `json:"id"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
	Method string `json:"method"`
}

var recipes = []Recipe{
	{
		ID: "0000000000000001",
		Ingredients: []Ingredient{
			{
				Name:   "corned beef brisket",
				Amount: "3 lb",
				Method: "whole",
			},
			{
				Name:   "spice pack",
				Amount: "1",
				Method: "",
			},
			{
				Name:   "cabbage",
				Amount: "1",
				Method: "quartered",
			},
			{
				Name:   "carrots",
				Amount: "1",
				Method: "1/2 bag",
			},
		},
	},
}
