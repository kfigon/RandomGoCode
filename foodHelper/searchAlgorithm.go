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

	strategy := getStrategy(includeStategyType)

	for _, v := range allFoods {
		commonIngredients := ingredients.intersection(v.requiredIngredients)
		fitness := calcFitness(commonIngredients, v.requiredIngredients)
		
		if strategy.shouldBeIncluded(ingredients, v) {
			f := v // go :(
			candidate := foodRecommendation{ f, fitness, }
			result = append(result, candidate)
		}
	}

	return result
}

type inclusionStrategy interface {
	shouldBeIncluded(usersIngredients *set, foodData food) bool
}

type fitnessInclusionStrategy struct {
	persentThreshold int
}

func (f fitnessInclusionStrategy) shouldBeIncluded(usersIngredients *set, foodData food) bool {
	commonIngredients := usersIngredients.intersection(foodData.requiredIngredients)
	fit := calcFitness(commonIngredients, foodData.requiredIngredients)
	return fit >= f.persentThreshold
}

func calcFitness(commonIngredients *set, required *set) int {
	return commonIngredients.size()/required.size() * 100
}

func getStrategy(strat includeStrategy) inclusionStrategy {
	switch strat {
	case defaultStrategy: return fitnessInclusionStrategy{ 100 }
	case eightyPercent: return fitnessInclusionStrategy{ 80 }
	}
	return fitnessInclusionStrategy{ 100 }
}

type includeStrategy string
const (
	defaultStrategy includeStrategy = "DEFAULT"
	eightyPercent = "80_PERCENT"
)