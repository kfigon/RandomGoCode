package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepareData() config {
	input := `
[foo]
key1 = 3
key2=some long name

[empty]

[x]
[another]
key1 = 1

[indented]
	asd = 123  `

	return parse(input)
}

func TestParseIni(t *testing.T) {
	data := prepareData()

	t.Run("present keys", func(t *testing.T) {
		assertPresent := func(section, property, expected string) {
			v, ok := data.get(section, property)
			assert.True(t, ok, fmt.Sprintf("%v:%v missing", section, property))
			assert.Equal(t, expected, v)
		}

		assertPresent("foo", "key1", "3")
		assertPresent("foo", "key2", "some long name")
		
		assertPresent("another", "key1", "1")
		
		assertPresent("indented", "asd", "123")
	})
	

	t.Run("ansent keys", func(t *testing.T) {
		assertAbsent := func(section, property string) {
			_, ok := data.get(section, property)
			assert.False(t, ok, fmt.Sprintf("%v:%v not expected, but found", section, property))
		}

		assertAbsent("foo", "empty")
		assertAbsent("x", "another")
	})
}

func TestLengths(t *testing.T) {
	data := prepareData()

	assert.Len(t, data, 5)
	assert.Len(t, data["foo"], 2)
	assert.Len(t, data["empty"], 0)
	assert.Len(t, data["x"], 0)
	assert.Len(t, data["another"], 1)
	assert.Len(t, data["indented"], 1)
}