package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type subscriber interface {
	update(string)
}
type publisher struct {
	susbcribers []subscriber
}

func (p *publisher) addSubscriber(sub subscriber) {
	p.susbcribers = append(p.susbcribers, sub)
}

func (p *publisher) notifySubscribers(state string) {
	for _, v := range p.susbcribers {
		v.update(state)
	}
}

func (p *publisher) businessLogic(someValue string) {
	// ... some logic
	p.notifySubscribers(someValue)
}

type subsciberCall func(string)
func (s subsciberCall) update(x string) {
	s(x)
}

func TestObserverPattern(t *testing.T) {
	var sub1 subsciberCall = func(x string) {
		assert.Equal(t, "foo", x)
	}
	var sub2 subsciberCall = func(x string) {
		assert.Equal(t, "foo", x)
	}

	p := publisher{}
	p.addSubscriber(sub1)
	p.addSubscriber(sub2)

	p.businessLogic("foo")
}