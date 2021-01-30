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

func newSearch(db productDb) *searchService {
	return &searchService{db}
}

func (s *searchService) findFoods(ingredients *set, includeStategyType includeStrategy) []food {
	result := make([]food,0)
	allFoods := s.db.findFoods()

	strategyFunction := getStrategy(includeStategyType)

	for _, v := range allFoods {		
		if shouldAdd(ingredients, v.requiredIngredients, strategyFunction) {
			f := v // go :(
			result = append(result, f)
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
func shouldAdd(ingredients *set, required *set, includeStrategy strategyFun) bool {
	commonIngredients := ingredients.intersection(required)
	return includeStrategy(ingredients.size(), required.size(), commonIngredients.size())
}