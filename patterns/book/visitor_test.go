package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// OOP approach to sum types
// decoupling behaviours, delegating to other classes. "Visiting" other classes to do something on them
// we can rely here on concrete classes + we have compile time guarantee to cover all cases

type petNamer interface {
	visit(visitor)
}

type dog string
func (d dog) visit(v visitor) {
	v.visitDog(d)
}

type cat string
func (c cat) visit(v visitor) {
	v.visitCat(c)
}

type visitor interface {
	visitDog(dog)
	visitCat(cat)
}

type visitorName struct {
	names []string
}

func (v *visitorName) visitDog(d dog) {
	v.names = append(v.names, string(d))
}

func (v *visitorName) visitCat(c cat) {
	v.names = append(v.names, string(c))
}

func TestVisitor(t *testing.T) {
	d1 := dog("fafik")
	d2 := dog("pufik")

	c1 := cat("mruczek")

	v := &visitorName{[]string{}}

	d1.visit(v)
	d2.visit(v)
	c1.visit(v)

	assert.Equal(t, []string{"fafik", "pufik", "mruczek"}, v.names)
}