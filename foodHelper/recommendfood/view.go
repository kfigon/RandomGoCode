package recommendfood

import (
	"html/template"
	"net/http"
)

func RecommendationView(writer http.ResponseWriter, results []FoodRecommendationDto) {
	tpl := template.Must(template.ParseFiles("recommendFood/results.html"))
	tpl.Execute(writer, results)
}

func RecommendationForm(writer http.ResponseWriter) {
	tpl := template.Must(template.ParseFiles("recommendFood/index.html"))
	tpl.Execute(writer, nil)
}
