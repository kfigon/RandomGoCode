package recommendfood

import (
	"net/http"
	"strings"
)

func GetNamesFromRequest(request *http.Request) []string {
	request.ParseForm()
	data := request.FormValue("data")
	return strings.Fields(data)
}
