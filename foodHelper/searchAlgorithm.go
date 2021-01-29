package main

type productDb interface{}

type searchService struct{
	db *productDb
}

type food struct{}

func newSearch(db productDb) *searchService {
	return &searchService{&db}
}

func (s *searchService) findFoods(ingredients *set) []food {
	return nil
}


