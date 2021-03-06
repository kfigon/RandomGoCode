package main

type foodDataProvider interface {
	findFoods() []food
}

type searchService struct {
	db foodDataProvider
}

func newSearch(db foodDataProvider) *searchService {
	return &searchService{db}
}

func (s *searchService) findFoods(ingredients *set, strategy inclusionStrategy) []foodRecommendation {
	result := make([]foodRecommendation, 0)
	allFoods := s.db.findFoods()

	for _, v := range allFoods {
		required := newSet(v.requiredIngredients...)
		if strategy.shouldBeIncluded(ingredients, required) {
			fitness := strategy.calcFitness(ingredients, required)

			f := v // go :(
			candidate := foodRecommendation{f, fitness}
			result = append(result, candidate)
		}
	}

	return result
}

func (s *searchService) findFoodsPercentageStrategy(ingredientIds []int) []foodRecommendation {
	ingredients := newSet(ingredientIds...)
	return s.findFoods(ingredients, fitnessInclusionStrategy{80})
}
