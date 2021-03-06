package recommentfood

type food struct {
	Name                string
	RequiredIngredients []int
}

type FoodRecommendation struct {
	food
	FitnessLevel int
}
