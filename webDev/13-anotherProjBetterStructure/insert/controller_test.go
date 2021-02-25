package getlist

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"mywebapp/model"
)

type insertFunction func(model.Element)

type mockDb struct {
	insertFun insertFunction
}

func (m mockDb) insert(el model.Element) {
	if m.insertFun != nil {
		m.insertFun(el)
	}
}

func TestInsert(t *testing.T) {
	db := mockDb{
		insertFun: func(el model.Element) {
			assert.Equal(t, 
				model.Element{Name:"Asd", Date: "some date"},
				el)
		},
	}
	controller := createInsertController(db)
	
	newElement := model.Element{Name:"Asd", Date: "some date"}
	err := controller.insert(newElement)
	assert.NoError(t, err)
}

func TestInsertInvalidElement(t *testing.T) {
	controller := createInsertController(mockDb{})
	
	invalidElement := model.Element{Date: "some date"}
	err := controller.insert(invalidElement)
	assert.NotNil(t, err)
}
