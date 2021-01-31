package main

import (
	"encoding/json"
	"testing"
	"strings"
)

// 1 - jajka
// 2 - makaron
// 3 - pomodory
// 4 - mielone
// 5 - kurczak
// 6 - salata
// 7 - oliwki
var foodJSON string = `[
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

var ingredientsJSON string = `[1,2,3,4,5,6,7]`

func TestReadFoodbase(t *testing.T) {
	provider := fromJson(strings.NewReader(foodJSON))
	foods := provider.findFoods()
	if len(foods) != 4 {
		t.Errorf("invalid food len, wanted: %v, got %v", 4, len(foods))
	}
	if foods[0].name != "spaghetti" {
		t.Errorf("invalid food name received: ", foods[0].name)
	}
	if foods[0].ingredients.size() != 3 {
		t.Error("Invalid ingredients received")
	}
}

