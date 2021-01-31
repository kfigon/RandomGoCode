package main

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
	Name string `json:"name"`
	Ingredients []int `json:"ingredients"`
}

func fromJSON(r io.Reader) foodDataProvider {
	dec := json.NewDecoder(r)
	var foodDto []foodJSONDto
	err := dec.Decode(&foodDto)
	if err != nil {
		return foodJSONProvider{}
	}
	var result foodJSONProvider
	result.foods = make([]food,0)
	for _, v := range foodDto {
		f := food{v.Name, newSet(v.Ingredients...)}
		result.foods = append(result.foods, f)
	}
	return result
}