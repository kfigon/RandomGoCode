package getlist

import (
	"mywebapp/model"
)

type GetListProvider interface {
	readList() []model.Element
}


type GetController struct {
	db GetListProvider
}

func CreateGetListController(db GetListProvider) *GetController {
	return &GetController{db}
}

func (c *GetController) GetList() []model.Element {
	return c.db.readList()
}

type MyDb struct {
	Elems []model.Element
}

func (m *MyDb) readList() []model.Element {
	return m.Elems
}

func MakeDb() *MyDb {
	return &MyDb{
		Elems: []model.Element {
			model.Element{Name: "foo", Date: "2021-02-25"},
			model.Element{Name: "bar", Date: "2021-02-30?!"},
		},
	}
}