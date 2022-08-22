package book

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// pattern to represent tree structures. Some methods dont have sense in leaf nodes!
type menu interface {
	description() string
	add(menu)
}

// leaf node
type menuItem struct {
	desc string
}

func (m *menuItem) description() string {
	return m.desc
}

func (m *menuItem) add(_ menu) {
	// no implementation
}

// composit
type subMenu struct {
	children []menu
	desc     string
}

func (s *subMenu) description() string {
	var childrenDesc []string
	for _, v := range s.children {
		childrenDesc = append(childrenDesc, v.description())
	}

	return s.desc + ": " + strings.Join(childrenDesc, ", ")
}

func (s *subMenu) add(child menu) {
	s.children = append(s.children, child)
}

func TestComposit(t *testing.T) {
	desserts := &subMenu{desc: "dessert menu", children: []menu{&menuItem{"chocolate"}, &menuItem{"ice cream"}}}
	dinnerMenu := &subMenu {
		desc: "dinner menu",
		children: []menu{
			&menuItem{"chicken"},
			&menuItem{"salad"},
			desserts,
		},
	}

	t.Run("simple flat menu", func(t *testing.T) {
		desc := desserts.description()
		assert.Equal(t, "dessert menu: chocolate, ice cream", desc)
	})

	t.Run("nested menu", func(t *testing.T) {
		desc := dinnerMenu.description()
		assert.Equal(t, "dinner menu: chicken, salad, dessert menu: chocolate, ice cream", desc)
	})

}
