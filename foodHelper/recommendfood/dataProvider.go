package recommendfood

import (
	"encoding/json"
	"io"
)

type foodJSONProvider struct {
	foods []food
}

func (f foodJSONProvider) findFoods() []food {
	return f.foods
}

type foodJSONDto struct {
	Name        string `json:"name"`
	Ingredients []int  `json:"ingredients"`
}

func FoodJSONProvider(r io.Reader) foodDataProvider {
	dec := json.NewDecoder(r)
	var foodDto []foodJSONDto
	err := dec.Decode(&foodDto)
	if err != nil {
		return foodJSONProvider{}
	}
	var result foodJSONProvider
	result.foods = make([]food, 0)
	for _, v := range foodDto {
		f := food{v.Name, v.Ingredients}
		result.foods = append(result.foods, f)
	}
	return result
}

type ingredientsDataProvider struct {
	ingredients []ingredient
}

func (i ingredientsDataProvider) getAll() []ingredient {
	return i.ingredients
}

type ingredientJSONDto struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func IngedientsJSONProvider(r io.Reader) ingredientsDataProvider {
	decoder := json.NewDecoder(r)
	var jsonData []ingredientJSONDto
	err := decoder.Decode(&jsonData)
	if err != nil {
		return ingredientsDataProvider{}
	}
	data := ingredientsDataProvider{}
	data.ingredients = make([]ingredient, 0)
	for _, v := range jsonData {
		data.ingredients = append(data.ingredients, ingredient{v.ID, v.Name})
	}
	return data
}
