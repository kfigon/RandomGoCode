package recommendfood

import "errors"

type ingredientProvider interface {
	getId(name string) (int, bool)
}

type recommendationController struct {
	ingredientDb  ingredientProvider
	searchService *searchService
}

func NewRecommendationController(ingDb ingredientProvider, service *searchService) *recommendationController {
	return &recommendationController{
		ingredientDb:  ingDb,
		searchService: service,
	}
}

func (r *recommendationController) FindFoods(names []string) []FoodRecommendationDto {
	return nil
}

const MAX_INGREDIENTS_NUM = 20
const MAX_INGREDIENT = 20

func validateIngredientNames(names []string) error {
	if ln := len(names); ln > MAX_INGREDIENTS_NUM || ln <= 0 {
		return errors.New("Invalid input length")
	}
	for _, v := range names {
		if ln := len(v); ln > MAX_INGREDIENT || ln <= 0 {
			return errors.New("Invalid element lenth")
		}
	}
	return nil
}
