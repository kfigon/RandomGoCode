package main

type food struct {
	name                string
	requiredIngredients []int
}

type foodRecommendation struct {
	food
	fitnessLevel int
}
