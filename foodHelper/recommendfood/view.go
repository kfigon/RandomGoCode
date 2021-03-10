package recommendfood

import (
	"html/template"
	"net/http"
)

func RecommendationView(writer http.ResponseWriter, results []FoodRecommendationDto) {
	tpl := template.Must(template.ParseFiles("results.html"))
	tpl.Execute(writer, results)
}

func RecommendationForm(writer http.ResponseWriter) {
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(writer, nil)
}
