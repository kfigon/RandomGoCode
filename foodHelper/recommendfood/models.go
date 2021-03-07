package recommendfood

type food struct {
	Name                string
	RequiredIngredients []int
}

type FoodRecommendation struct {
	food
	FitnessLevel int
}

type ingredient struct {
	ID   int
	Name string
}
