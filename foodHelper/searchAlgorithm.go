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

func (s *searchService) findFoods(ingredients *set, strategy inclusionStrategy) []foodRecommendation {
	result := make([]foodRecommendation,0)
	allFoods := s.db.findFoods()

	for _, v := range allFoods {

		if strategy.shouldBeIncluded(ingredients, v.requiredIngredients) {
			f := v // go :(
			fitness := strategy.calcFitness(ingredients, v.requiredIngredients)
			candidate := foodRecommendation{ f, fitness, }
			result = append(result, candidate)
		}
	}

	return result
}