package getlist

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"mywebapp/model"
	"testing"
)

func TestRenderThings(t *testing.T) {
	elems := []model.Element {
		model.Element{Name: "first", Date: "foo"},
		model.Element{Name: "second", Date: "bar"},
	}
	var writer bytes.Buffer
	Render(&writer, elems)
	result := writer.String()

	assert.Contains(t, result, "To do list")
	
	assert.Contains(t, result, "first")
	assert.Contains(t, result, "foo")

	assert.Contains(t, result, "second")
	assert.Contains(t, result, "bar")
}