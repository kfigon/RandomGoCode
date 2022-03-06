package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {

	t := initTrie()
	http.HandleFunc("/api/healthcheck", healthcheck)
	http.HandleFunc("/api/ingredients", handleIngredients(t))
	http.HandleFunc("/api/suggestions", handleSuggestions(foodProvider()))

	port := 8080
	log.Println("Starting on port", port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func initTrie() *trie {
	// todo: fill with data
	return &trie{}
}

func foodProvider() dataProvider {
	// todo
	return nil
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	out := struct {
		Status string `json:"status"`
	}{"ok"}
	toJson(w, &out)
}

func toJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func handleIngredients(t *trie) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			invalidMethod(w)
			return
		}
		pref := r.FormValue("prefix")
		res := t.suggestions(pref)

		out := struct {
			Ingredients []string `json:"ingredients"`
		}{res}
		toJson(w, &out)
	}
}

func invalidMethod(w http.ResponseWriter) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

type suggestion struct {
	name string
	description string
	ingredients []string
}

type dataProvider interface {
	allSuggestions() []suggestion
}
func handleSuggestions(db dataProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			invalidMethod(w)
			return
		}
		
		type request struct {
			Ingredients []string `json:"ingredients"`
		}
		var reqBody request
		json.NewDecoder(r.Body).Decode(&reqBody)

		suggestions := findSuggestions(db.allSuggestions(), reqBody.Ingredients)
		
		type result struct {
			Name string `json:"Name"`
			Description string `json:"Description"`
			Ingredients []string `json:"Ingredients"`
		}
		type response struct {
			Results []result `json:"results"`
		}

		var out response
		for _, s := range suggestions {
			out.Results = append(out.Results, result{
				Name: s.name,
				Description: s.description,
				Ingredients: s.ingredients,
			})
		}
		toJson(w, &out)
	}
}

type void struct{}
type set map[string]void

func findSuggestions(allSuggestions []suggestion, givenIngredients []string) []suggestion {
	givenSet := buildSet(givenIngredients)

	var out []suggestion
	for _, v := range allSuggestions {
		res := match(v, givenSet)
		if res >= 70 {
			out = append(out, v)
		}
	}
	return out
}

func buildSet(ing []string) set {
	s := set{}
	for _, v := range ing {
		s[v] = void{}
	}
	return s
}

func match(s suggestion, givenSet set) int {
	allIngSet := buildSet(s.ingredients)
	matched := 0
	for k := range allIngSet {
		if _, ok := givenSet[k]; ok {
			matched++
		}
	}

	return (100*matched)/len(allIngSet)
}