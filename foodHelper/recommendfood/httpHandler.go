package recommendfood

import (
	"log"
	"net/http"
	"strings"
)

func GetNamesFromRequest(request *http.Request) []string {
	request.ParseForm()
	data := request.FormValue("data")
	log.Println("Provided: ", data)
	return strings.Fields(data)
}
