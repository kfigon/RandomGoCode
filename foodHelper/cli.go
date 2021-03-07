package main

import (
	"flag"
	"fmt"
	"foodHelper/recommendfood"
	"strconv"
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

var knownIngredients = map[string]int{
	"jajka":    1,
	"makaron":  2,
	"pomodory": 3,
	"mielone":  4,
	"kurczak":  5,
	"salata":   6,
	"oliwki":   7,
}

func readKeyFromValue(val int) string {
	for key := range knownIngredients {
		if knownIngredients[key] == val {
			return key
		}
	}
	return ""
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	foodProvider := recommendfood.FoodJSONProvider(strings.NewReader(inputJson))
	search := recommendfood.NewSearch(foodProvider)

	var threshold = flag.Int("prog", 100, "Prog dolaczenia jedzenia <0;100>")
	var foods arrayFlags
	flag.Var(&foods, "f", "idx of ingredient")

	flag.Parse()

	userFoods := parseFoods(foods)
	foodRecommendations := search.RecommendFoods(userFoods, *threshold)

	printKnownIngredients()
	fmt.Println("\nProvided:")
	fmt.Println(ingredientsString(userFoods))
	fmt.Println("---------")
	for _, v := range foodRecommendations {
		printFoodRecomendation(v)
	}
}

func parseFoods(foods arrayFlags) []int {
	var out []int
	for _, v := range foods {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		if readKeyFromValue(i) != "" {
			out = append(out, i)
		}
	}
	return out
}

func printFoodRecomendation(v recommendfood.FoodRecommendation) {
	fmt.Printf("%v, fit: %v, ingredients: %v\n", v.Name, v.FitnessLevel, v.RequiredIngredients)
}

func ingredientsString(ing []int) string {
	out := ""
	for _, v := range ing {
		out += readKeyFromValue(v) + " "
	}
	return out
}

func printKnownIngredients() {
	fmt.Println("Known ingredients")
	for key := range knownIngredients {
		fmt.Printf("%v -> %v\n", key, knownIngredients[key])
	}
}
