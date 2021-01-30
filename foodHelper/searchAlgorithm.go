package main

type productDb interface {
	findFoods() []food
}

type searchService struct {
	db *productDb
}

type food struct {
	name string
	requiredIngredients *set
}

func newSearch(db productDb) *searchService {
	return &searchService{&db}
}

func (s *searchService) findFoods(ingredients *set) []food {
	return nil
}


