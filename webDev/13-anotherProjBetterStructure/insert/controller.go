package getlist

import (
	"mywebapp/model"
)

type insertHandler interface {
	insert(el model.Element)
}


type insertController struct {
	db insertHandler
}

func createInsertController(db insertHandler) *insertController {
	return &insertController{db}
}

func (c *insertController) insert(element model.Element) {
	c.db.insert(element)
}