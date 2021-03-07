package recommendfood

import "strings"

func createMockFoodDb() foodDataProvider {
	return FoodJSONProvider(strings.NewReader(foods))
}

func createMockIngredientsDb() ingredientProvider {
	return IngedientsJSONProvider(strings.NewReader(ingredients))
}

// language=json
var ingredients string = `[
	{ "id": 0, "name": "egg" },
	{ "id": 1, "name": "chicken" },
	{ "id": 2, "name": "beef" },
	{ "id": 3, "name": "salmon" },
	{ "id": 4, "name": "salad" },
	{ "id": 5, "name": "cheese" },
	{ "id": 6, "name": "apple" },
	{ "id": 7, "name": "noodle" },
	{ "id": 8, "name": "bread" },
	{ "id": 9, "name": "tomato" },
	{ "id": 10, "name": "feta" },
	{ "id": 11, "name": "cucumber" }
]`

// language=json
var foods string = `[
	{"name":  "first", "ingredients":[0,1,3]},
	{"name":  "second", "ingredients":[0,1,4]},
	{"name":  "third", "ingredients":[4,5,6]}
]`
