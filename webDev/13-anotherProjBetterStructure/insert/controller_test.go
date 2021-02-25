package insert

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"mywebapp/model"
)

type insertFunction func(model.Element)

type mockDb struct {
	insertFun insertFunction
}

func (m mockDb) Insert(el model.Element) {
	if m.insertFun != nil {
		m.insertFun(el)
	}
}

func TestInsert(t *testing.T) {
	db := &mockDb{
		insertFun: func(el model.Element) {
			assert.Equal(t, 
				model.Element{Name:"Asd", Date: "some date"},
				el)
		},
	}
	controller := CreateInsertController(db)
	
	newElement := model.Element{Name:"Asd", Date: "some date"}
	err := controller.Insert(newElement)
	assert.NoError(t, err)
}

func TestInsertInvalidElement(t *testing.T) {
	controller := CreateInsertController(mockDb{})
	
	invalidElement := model.Element{Date: "some date"}
	err := controller.Insert(invalidElement)
	assert.NotNil(t, err)
}
