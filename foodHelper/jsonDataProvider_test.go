package main

import (
	"encoding/json"
	"testing"
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
	data := `[
		{}
		]`
}

