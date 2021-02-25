package getlist

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"mywebapp/model"
)

type mockDb struct {
	elems []model.Element
}

func (m mockDb) readList() [] model.Element{
	return m.elems
}

func TestReadEmptyList(t *testing.T) {
	controller := CreateGetListController(mockDb{})
	list := controller.GetList()
	assert.Empty(t, list)
}

func TestReadNotEmptyList(t *testing.T) {
	elemenets := []model.Element{
		model.Element{Name: "First task", Date: "2021-02-25"},
	}
	controller := CreateGetListController(mockDb{elemenets})
	list := controller.GetList()
	assert.Contains(t, list, model.Element{Name: "First task", Date: "2021-02-25"})
}