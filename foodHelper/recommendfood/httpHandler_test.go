package recommendfood

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExtractNames(t *testing.T) {
	tdt := []struct {
		desc        string
		requestData string
		expected    []string
	}{
		{"Empty data in request", "", []string{}},
		{"Single data", "data=asd", []string{"asd"}},
		{"More data", "data=asd foo bar", []string{"asd", "foo", "bar"}},
	}
	for _, tc := range tdt {
		t.Run(tc.desc, func(t *testing.T) {
			request := httptest.NewRequest("POST", "http://localhost:8080", strings.NewReader(tc.requestData))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			result := GetNamesFromRequest(request)
			assert.Equal(t, tc.expected, result)
		})
	}
}
