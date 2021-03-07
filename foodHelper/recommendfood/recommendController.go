package recommendfood

import "errors"

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

type foundState int

const (
	NOT_FOUND foundState = iota
	FOUND
	ADJUSTED
)

type adjustResult struct {
	foundName   string
	foundId     int
	matchResult foundState
}

func (r *recommendationController) adjustName(input string) adjustResult {
	MAX_HAMMING_DISTANCE := 2

	allIngredients := r.ingredientDb.getAll()
	bestGuess := struct {
		distance int
		id       int
		name     string
	}{}
	bestGuess.distance = 999999999999

	for _, ing := range allIngredients {
		if input == ing.Name {
			return adjustResult{
				foundName:   ing.Name,
				foundId:     ing.ID,
				matchResult: FOUND,
			}
		}
		if ham := calcHammingDistance(input, ing.Name); ham < bestGuess.distance {
			bestGuess.distance = ham
			bestGuess.name = ing.Name
			bestGuess.id = ing.ID
		}
	}

	if bestGuess.distance < MAX_HAMMING_DISTANCE {
		return adjustResult{
			foundName:   bestGuess.name,
			foundId:     bestGuess.id,
			matchResult: ADJUSTED,
		}
	}

	return adjustResult{matchResult: NOT_FOUND}
}

func calcHammingDistance(a string, b string) int {
	return -1
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
