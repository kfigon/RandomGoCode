package recommendfood

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestExtractNames(t *testing.T) {
	tdt := []struct {
		desc     string
		req      io.Reader
		expected []string
	}{
		{"Empty data in request", nil, []string{}},
	}
	for _, tc := range tdt {
		t.Run(tc.desc, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/", tc.req)
			result := GetNamesFromRequest(request)
			assert.Equal(t, tc.expected, result)
		})
	}
}
