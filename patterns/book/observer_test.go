package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// observer - notify multiple clients about some state change
// decompose clients from main source of change

type weatherStation struct {
	// or functional - just callbacks
	subscribers []subscriber
}

func (w *weatherStation) addSubscriber(sub subscriber) {
	w.subscribers = append(w.subscribers, sub)
}

// loose coupling. Other way - exlpicit call on each subscriber. No extensibility, new services require change in weather station also
// displayService.notify()
// statistic.notify()
func (w *weatherStation) changeState(newTemperature int) {
	for _, v := range w.subscribers {
		v.notify(newTemperature)
	}
}

// we can work in PULL or PUSH mode. Push here
// pull - send a reference to weather station to subscriber, it cal pull all data he wants (notify(*weatherStation))
// push - weather station pushes data directly as an argument
type subscriber interface {
	notify(int)
}

type statisticsService int
func (s *statisticsService) notify(newVal int) {
	*s = statisticsService(newVal)
}

type displayService int
func (d *displayService) notify(newVal int) {
	*d = displayService(newVal)
}

func TestObserverPattern(t *testing.T) {
	w := &weatherStation{}
	stat := statisticsService(0)
	disp := displayService(0)

	w.addSubscriber(&stat)
	w.addSubscriber(&disp)

	assert.Equal(t, 0, int(stat))
	assert.Equal(t, 0, int(stat))

	w.changeState(123)

	assert.Equal(t, 123, int(stat))
	assert.Equal(t, 123, int(stat))
}
