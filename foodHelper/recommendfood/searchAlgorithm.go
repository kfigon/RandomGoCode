package recommendfood

type foodDataProvider interface {
	findFoods() []food
}

type searchService struct {
	db foodDataProvider
}

func NewSearch(db foodDataProvider) *searchService {
	return &searchService{db}
}

func (s *searchService) findFoods(ingredients *set, strategy inclusionStrategy) []FoodRecommendation {
	result := make([]FoodRecommendation, 0)
	allFoods := s.db.findFoods()

	for _, v := range allFoods {
		required := newSet(v.RequiredIngredients...)
		if strategy.shouldBeIncluded(ingredients, required) {
			fitness := strategy.calcFitness(ingredients, required)

			f := v // go :(
			candidate := FoodRecommendation{f, fitness}
			result = append(result, candidate)
		}
	}

	return result
}

func (s *searchService) RecommendFoods(ingredients []int, inclusionPercentage int) []FoodRecommendation {
	return s.findFoods(newSet(ingredients...), fitnessInclusionStrategy{inclusionPercentage})
}
