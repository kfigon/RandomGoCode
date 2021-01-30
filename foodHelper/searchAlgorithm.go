package main

type productDb interface {
	findFoods() []food
}

type searchService struct {
	db productDb
}

type food struct {
	name string
	requiredIngredients *set
}

type foodRecommendation struct {
	food
	fitnessLevel int
}

func newSearch(db productDb) *searchService {
	return &searchService{db}
}

func (s *searchService) findFoods(ingredients *set, includeStategyType includeStrategy) []foodRecommendation {
	result := make([]foodRecommendation,0)
	allFoods := s.db.findFoods()

	strategyFunction := getStrategy(includeStategyType)

	for _, v := range allFoods {		
		commonIngredients := ingredients.intersection(v.requiredIngredients)
		
		if shouldAdd(ingredients, v.requiredIngredients, commonIngredients, strategyFunction) {
			fitness := calcFitness(ingredients, v.requiredIngredients)
			f := v // go :(
			candidate := foodRecommendation{ f, fitness, }
			result = append(result, candidate)
		}
	}

	return result
}

func getStrategy(strat includeStrategy) strategyFun {
	if strat == defaultStrategy {
		return defaultIncludeStrategy
	}
	return defaultIncludeStrategy
}

type includeStrategy string
const (
	defaultStrategy includeStrategy = "DEFAULT"
)

type strategyFun func(ingredientSize, requiredSize, commonIngredientSize int) bool

// all required provided
func defaultIncludeStrategy(ingredientSize, requiredSize, commonIngredientSize int) bool{
	return requiredSize == commonIngredientSize
}
func shouldAdd(ingredients *set, required *set, commonIngredients *set, includeStrategy strategyFun) bool {
	return includeStrategy(ingredients.size(), required.size(), commonIngredients.size())
}

func calcFitness(ingredients *set, required *set) int {
	return ingredients.size()/required.size() * 100
}