package main

import (
	// "log"
)

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

func newSearch(db productDb) *searchService {
	return &searchService{db}
}

func (s *searchService) findFoods(ingredients *set) []food {
	result := make([]food,0)
	allFoods := s.db.findFoods()

	for _, v := range allFoods {
		commonIngredients := ingredients.intersection(v.requiredIngredients)

		// log.Printf("Intesection of %v with %v, got: %v\n", ingredients.els(), v.requiredIngredients.els(), commonIngredients.els())

		if commonIngredients.size() != 0 {
			f := v
			result = append(result, f)
		}
	}

	return result
}


