package main

import (
	"fmt"
	"strings"
)

var inputJson string = `[
    {
        "name": "spaghetti",
        "ingredients": [2,3,4]
    },
    {
        "name": "jajecznica",
        "ingredients": [1]
    },
    {
        "name": "salatka z jajkiem",
        "ingredients": [6,7,1]
    },
    {
        "name": "salatka z kurczakiem",
        "ingredients": [6,7,5]
    }
]`

var knownIngredients = map[string]int {
	"jajka" : 1,
	"makaron" : 2,
	"pomodory" : 3,
	"mielone" : 4,
	"kurczak" : 5,
	"salata" : 6,
	"oliwki" : 7,
}

func readKeyFromValue(val int) string {
	for key := range knownIngredients {
		if knownIngredients[key] == val {
			return key
		}
	}
	return ""
}


func main() {
	foodProvider := fromJSON(strings.NewReader(inputJson))
	search := newSearch(foodProvider)

	threshold := 100
	userFoods := newSet(5,4,1)
	foodRecomendations := search.findFoods(userFoods, fitnessInclusionStrategy{threshold})

	
	fmt.Println(ingredientsString(userFoods))
	for _, v := range foodRecomendations {
		printFoodRecomendation(v)
	}
}

func printFoodRecomendation(v foodRecommendation) {
	fmt.Printf("%v, fit: %v, ingredients: %v\n", v.name, v.fitnessLevel, ingredientsString(v.requiredIngredients))
}

func ingredientsString(ing *set) string {
	out := ""
	for _,v := range ing.els() {
		out += readKeyFromValue(v)+" "
	}
	return out
}
