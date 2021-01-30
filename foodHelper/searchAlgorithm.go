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
		fitness := calcFitness(commonIngredients, v.requiredIngredients)
		
		if strategyFunction(fitness) {
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

type strategyFun func(fitnessLevel int) bool

// all required provided
func defaultIncludeStrategy(fitnessLevel int) bool{
	return fitnessLevel == 100
}

func calcFitness(commonIngredients *set, required *set) int {
	return commonIngredients.size()/required.size() * 100
}