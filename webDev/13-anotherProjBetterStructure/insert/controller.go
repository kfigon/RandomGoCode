package insert

import (
	"fmt"
	"mywebapp/model"
)

type InsertHandler interface {
	Insert(el model.Element)
}


type InsertController struct {
	db InsertHandler
}

func CreateInsertController(db InsertHandler) *InsertController {
	return &InsertController{db}
}

func (c *InsertController) Insert(element model.Element) error {
	if len(element.Name) == 0 || len(element.Name) > 100 {
		return fmt.Errorf("Invalid element given")
	}
	c.db.Insert(element)
	return nil
}