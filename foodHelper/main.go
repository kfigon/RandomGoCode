package main

import (
	"foodHelper/recommendfood"
	"log"
	"net/http"
)

func main() {

	controller := recommendfood.NewRecommendationController(recommendfood.CreateMockIngredientsDb(), recommendfood.NewSearch(recommendfood.CreateMockFoodDb()))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		recommendfood.RecommendationForm(writer)
	})

	mux.HandleFunc("/recommend", func(writer http.ResponseWriter, request *http.Request) {
		names := recommendfood.GetNamesFromRequest(request)
		results := controller.FindFoods(names)
		recommendfood.RecommendationView(writer, results)
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
