package getlist

import (
	"fmt"
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

func (c *insertController) insert(element model.Element) error {
	if len(element.Name) == 0 || len(element.Name) > 100 {
		return fmt.Errorf("Invalid element given")
	}
	c.db.insert(element)
	return nil
}