package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {

	data := foodProvider()
	t := initTrie(data)

	http.Handle("/", http.FileServer(http.Dir("./client/public")))
	http.HandleFunc("/api/healthcheck", healthcheck)
	http.HandleFunc("/api/ingredients", handleIngredients(t))
	http.HandleFunc("/api/suggestions", handleSuggestions(data))

	port := 8000
	log.Println("Starting on port", port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func initTrie(data dataProvider) *trie {
	t := &trie{}
	for _, v := range data.allSuggestions() {
		for _, ing := range v.ingredients {
			t.add(ing)
		}
	}
	return t
}

type inmemoryDataProvider func() []suggestion
func (idp inmemoryDataProvider) allSuggestions() []suggestion {
	return idp()
}

func foodProvider() dataProvider {
	// todo: fill data and images
	results := []suggestion {
		{
			name: "Jajecznica",
			description: "Smaczna lol",
			ingredients: []string{"jajko", "chlebek", "maslo"},
		},
		{
			name: "Jajko sadzone",
			description: "tasty",
			ingredients: []string{"jajko", "chlebek", "olej"},
		},
		{
			name: "Kurczak po chinsku",
			description: "wow",
			ingredients: []string{"kurczak", "makaron", "warzywa", "grzyby mun", "sos sojowy"},
		},
	}
	f := func() []suggestion {
		return results
	}
	return inmemoryDataProvider(f)
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
	img []byte
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
			Image string `json:"image"`
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
				Image: encodeImage(s.img),
			})
		}
		toJson(w, &out)
	}
}

func encodeImage(data []byte) string {
	return "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(data)
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