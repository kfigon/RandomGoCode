package day5

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/8

func readFile(t *testing.T) []int {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)

	strContent := string(content)
	out := make([]int, 0)
	for i := 0; i < len(strContent); i++ {
		v, err := strconv.Atoi(string(strContent[i]))
		require.NoError(t, err)
		
		out = append(out, v)
	}
	return out
	
}
func TestParsingLayers(t *testing.T) {
	in := []int{1,2,3,4,5,6,7,8,9,0,1,2}
	width,height := 3,2

	image := parseInput(in, width, height)

	assert.Equal(t, 2, image.numberOfLayers())
	assert.Equal(t, []int{1,2,3}, image.extractRow(0,0))
	assert.Equal(t, []int{4,5,6}, image.extractRow(0,1))
	assert.Equal(t, []int{7,8,9}, image.extractRow(1,0))
	assert.Equal(t, []int{0,1,2}, image.extractRow(1,1))
}

func TestParsingLayers2(t *testing.T) {
	in := []int{1,2,3,4,5,6,7,8,9,0,1,2}
	width,height := 3,2

	image := parseInput(in, width, height)

	assert.Equal(t, 2, image.numberOfLayers())
	assert.Equal(t, []int{1,2,3,4,5,6}, image.extractLayer(0))
	assert.Equal(t, []int{7,8,9,0,1,2}, image.extractLayer(1))
}

func TestPart1(t *testing.T) {
	file := readFile(t)
	const width int = 25
	const height int = 6

	image := parseInput(file, width, height)
	layer := image.findLayerWithSmallestNumberOf(0)
	num := countNumberOfMultiplied(layer, 1,2)

	assert.Equal(t, 123,num)
}

func countNumberOfMultiplied(layerData []int, first, second int) int {
	return -1
}

type size struct {
	width int
	height int
}

func parseInput(data []int, width, height int) *image {
	return &image {
		in: data,
		s: size{width: width,height: height},
	}
}

type image struct {
	in []int
	s size
}

func (i *image) numberOfLayers() int { return len(i.in)/i.layerSize() }
func (i *image) layerSize() int { return i.s.width*i.s.height }

func (i *image) extractRow(layerNum, rowNum int) []int {
	if layerNum >= i.numberOfLayers() || rowNum >= i.s.height {
		return []int{}
	}

	start := layerNum * i.layerSize() + rowNum*i.s.width
	stop := start+i.s.width
	return i.in[start:stop]
}

func (i *image) extractLayer(layerNum int) []int {
	if layerNum >= i.numberOfLayers(){
		return []int{}
	}

	start := layerNum * i.layerSize()
	stop := start+i.layerSize()
	return i.in[start:stop]
}

func (i *image) findLayerWithSmallestNumberOf(numToFind int) []int {
	return []int{}
}