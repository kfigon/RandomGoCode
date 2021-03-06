package main

type foodDataProvider interface {
	findFoods() []food
}

type searchService struct {
	db foodDataProvider
}

type food struct {
	name string
	requiredIngredients *set
}

type foodRecommendation struct {
	food
	fitnessLevel int
}

func newSearch(db foodDataProvider) *searchService {
	return &searchService{db}
}

func (s *searchService) findFoods(ingredients *set, strategy inclusionStrategy) []foodRecommendation {
	result := make([]foodRecommendation,0)
	allFoods := s.db.findFoods()

	for _, v := range allFoods {

		if strategy.shouldBeIncluded(ingredients, v.requiredIngredients) {
			fitness := strategy.calcFitness(ingredients, v.requiredIngredients)
			
			f := v // go :(
			candidate := foodRecommendation{ f, fitness, }
			result = append(result, candidate)
		}
	}

	return result
}

func (s *searchService) findFoodsPercentageStrategy(ingredientIds []int) []foodRecommendation {
	ingredients := newSet(ingredientIds...)
	return s.findFoods(ingredients, fitnessInclusionStrategy{})
}