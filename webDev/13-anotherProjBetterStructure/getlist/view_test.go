package getlist

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"mywebapp/model"
	"testing"
)
func init() {
	BASE_PATH = "../"
}

func TestRenderThings(t *testing.T) {
	elems := []model.Element {
		model.Element{Name: "first", Date: "foo"},
		model.Element{Name: "second", Date: "bar"},
	}
	var writer bytes.Buffer

	Render(&writer, elems)
	result := writer.String()

	assert.Contains(t, result, "To do list", "header not present")
	
	assert.Contains(t, result, "first", "first elem not present")
	assert.Contains(t, result, "foo", "first elem not present")

	assert.Contains(t, result, "second", "second elem not present")
	assert.Contains(t, result, "bar", "second elem not present")
}