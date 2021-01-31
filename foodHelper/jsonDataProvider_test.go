package main

import (
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

func TestReadFoodbase(t *testing.T) {
	provider := fromJSON(strings.NewReader(foodJSON))
	foods := provider.findFoods()
	if len(foods) != 4 {
		t.Errorf("invalid food len, wanted: %v, got %v", 4, len(foods))
	}
	if foods[0].name != "spaghetti" {
		t.Errorf("invalid food name received: %v", foods[0].name)
	}
	if foods[0].requiredIngredients.size() != 3 {
		t.Error("Invalid ingredients received")
	}
}

func TestReadInvalidFoodbase(t *testing.T) {
	provider := fromJSON(strings.NewReader("[{asd}]"))
	foods := provider.findFoods()
	if len(foods) != 0 {
		t.Errorf("invalid food len, wanted: %v, got %v", 0, len(foods))
	}
}
