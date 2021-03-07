package recommendfood

import "errors"

type foundState int

const (
	FOUND = iota
	NOT_FOUND
	GUESS
)

type ingredientProvider interface {
	getAll() []ingredient
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

var tooMuchInputError = errors.New("TOO_MUCH_INPUT_ERROR")
var tooBigIngredient = errors.New("TOO_BIG_INGREDIENT")

func validateIngredientNames(names []string) error {

	if ln := len(names); ln > MAX_INGREDIENTS_NUM || ln <= 0 {
		return tooMuchInputError
	}
	for _, v := range names {
		if ln := len(v); ln > MAX_INGREDIENT || ln <= 0 {
			return tooBigIngredient
		}
	}
	return nil
}
