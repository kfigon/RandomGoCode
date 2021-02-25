package landing

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	BASE_PATH = "../"
}

func TestRenderLandingPage(t *testing.T) {
	var writer bytes.Buffer

	Render(&writer)
	result := writer.String()

	assert.Contains(t, result, "hi there ziomx")
	assert.Contains(t, result, "Lorem ipsum")
}