package getlist

import (
	"mywebapp/model"
)

type getListProvider interface {
	readList() []model.Element
}


type getController struct {
	db getListProvider
}

func createGetListController(db getListProvider) *getController {
	return &getController{db}
}

func (c *getController) getList() []model.Element {
	return c.db.readList()
}