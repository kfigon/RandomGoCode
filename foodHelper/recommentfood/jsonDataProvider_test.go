package recommentfood

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// 1 - jajka
// 2 - makaron
// 3 - pomodory
// 4 - mielone
// 5 - kurczak
// 6 - salata
// 7 - oliwki
var foodJSON string = `[
    {
        "name": "spaghetti",
        "ingredients": [2,3,4]
    },
    {
        "name": "jajecznica",
        "ingredients": [1]
    },
    {
        "name": "salatka z jajkiem",
        "ingredients": [6,7,1]
    },
    {
        "name": "salatka z kurczakiem",
        "ingredients": [6,7,5]
    }
]`

func TestReadFoodbase(t *testing.T) {
	provider := FromJSON(strings.NewReader(foodJSON))
	foods := provider.findFoods()
	assert.Equal(t, 4, len(foods))
	assert.Equal(t, "spaghetti", foods[0].Name)
	assert.Equal(t, 3, len(foods[0].RequiredIngredients))
}

func TestReadInvalidFoodbase(t *testing.T) {
	provider := FromJSON(strings.NewReader("[{asd}]"))
	foods := provider.findFoods()
	assert.Equal(t, 0, len(foods))
}
