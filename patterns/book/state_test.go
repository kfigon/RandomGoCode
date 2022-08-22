package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
// state machines make easy - without huge switch statement and ifs
// makes state transitions and adding new states easy

// methods - available inputs
// implementations - available states
type state interface {
	insertCoin()
	getGum()
	getMoneyBack()
}

type noMoneyState struct{
	*gumMachine
}

func (n *noMoneyState) insertCoin(){
	n.gumMachine.currentState = n.gumMachine.moneyPresent
	n.gumMachine.stateName = "coin"
}
func (n *noMoneyState) getGum(){
	n.gumMachine.stateName = "no money"
}
func (n *noMoneyState) getMoneyBack(){
	n.gumMachine.stateName = "no money"
	n.gumMachine.currentState = n.gumMachine.noMoney
}

type moneyPresentState struct{
	*gumMachine
}

func (m *moneyPresentState) insertCoin(){
	m.gumMachine.currentState = m.gumMachine.moneyPresent
	m.gumMachine.stateName = "coin"
}
func (m *moneyPresentState) getGum(){
	m.gumMachine.currentState = m.gumMachine.noMoney
	m.gumMachine.stateName = "gum"
}
func (m *moneyPresentState) getMoneyBack(){
	m.gumMachine.currentState = m.gumMachine.noMoney
	m.gumMachine.stateName = "no money"
}

type gumMachine struct{
	stateName string
	currentState state
	noMoney state
	moneyPresent state
}

func (g *gumMachine) insertCoin() {
	g.currentState.insertCoin()
}

func (g *gumMachine) getGum() {
	g.currentState.getGum()
}

func (g *gumMachine) getMoneyBack() {
	g.currentState.getMoneyBack()
}

func newMachine() *gumMachine {
	this := &gumMachine{
		stateName: "no money",
	}
	this.currentState = &noMoneyState{this}
	this.noMoney = &noMoneyState{this}
	this.moneyPresent = &moneyPresentState{this}

	return this
}

func TestState(t *testing.T) {
	machine := newMachine()

	assert.Equal(t, "no money", machine.stateName)
	machine.insertCoin()
	assert.Equal(t, "coin", machine.stateName)
	machine.getGum()
	assert.Equal(t, "gum", machine.stateName)
}