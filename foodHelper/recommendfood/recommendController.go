package recommendfood

type ingredientProvider interface{}

type recommendationController struct {
	ingredientDb  ingredientProvider
	searchService searchService
}

func NewRecommendationController(ingDb ingredientProvider, service searchService) *recommendationController {
	return &recommendationController{
		ingredientDb:  ingDb,
		searchService: service,
	}
}
