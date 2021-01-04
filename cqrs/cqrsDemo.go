package main

import (
	"strings"
	"fmt"
	"errors"
)

// Command Query Responsibility Segregation
// command - object that changes app state, no return
// query - fetch data, no state change

type PersonModel struct {
	Id int
	Name string
	Age uint
	Street string
}

type ReadQuery struct {
	userId int
}

type CreateCommand struct {
	Name string
	Age uint
	Street string
}

type UpdateCommand struct {
	Name string
	Age uint
	Street string
	Id int
}

type DB map[int]PersonModel

type Application struct {
	db DB
}

// no change in state
func (app *Application) readHandler(cmd ReadQuery) (PersonModel, error) {
	val, ok := app.db[cmd.userId]
	if !ok {
		return val, errors.New(fmt.Sprint("No object with id ", cmd.userId))
	}
	return val, nil
}

// return void
func (app *Application) createUser(cmd CreateCommand) error {
	if len(strings.TrimSpace(cmd.Name)) == 0 {
		return errors.New("Empty name provided")
	} else if len(strings.TrimSpace(cmd.Street)) == 0 {
		return errors.New("Empty street provided")
	}
	newID := len(app.db)
	newPerson := PersonModel{
		Age: cmd.Age,
		Name: cmd.Name,
		Street: cmd.Street,
		Id: newID,
	}
	app.db[newID] = newPerson
	return nil
}

func (app *Application) updatePerson(cmd UpdateCommand) error {
	val, ok := app.db[cmd.Id]
	if !ok {
		return errors.New(fmt.Sprint("No user with id", cmd.Id))
	}

	val.Age = cmd.Age
	val.Name = cmd.Name
	val.Street = cmd.Street

	app.db[cmd.Id] = val
	return nil
}

func read(app Application, cmd ReadQuery) {
	usr, err := app.readHandler(cmd)
	if err != nil {
		fmt.Println("Got error during reading:", err.Error())
	} else {
		fmt.Println("Read user", usr)
	}
}

func main() {
	db := make(map[int]PersonModel)
	app := Application{db: db}
	app.createUser(CreateCommand{Age: 5, Name: "Foo", Street:"Bars"})
	app.createUser(CreateCommand{Age: 15, Name: "Asd", Street:"xax"})
	
	err := app.createUser(CreateCommand{Name: "", Street:""})
	if err == nil {
		fmt.Println("should have error here!")
	}

	read(app, ReadQuery{userId: 1})
	
	err = app.updatePerson(UpdateCommand{Id: 1, Name: "Name changed!"})
	if err != nil {
		fmt.Println("Got error during update")
	}

	read(app, ReadQuery{userId: 1})
}


