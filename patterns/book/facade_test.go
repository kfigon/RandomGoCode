package book

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// aggregate multiple classes (with composition) and simplify the interface

type lightController struct{}
func (l *lightController) light() string {
	return "light on"
}

type musicController struct{}
func (m *musicController) playMusic() string {
	return "play music"
}

type videoController struct{}
func (v *videoController) playSomeVideo() string {
	return "play video"
}

type partyFacade struct {
	l *lightController
	m *musicController
	v *videoController
}

// just wrap around multiple components and expose simplify interface
func (p *partyFacade) partyOn() string {
	return fmt.Sprintf("%s %s %s", p.l.light(), p.m.playMusic(), p.v.playSomeVideo())
}

func TestFacade(t *testing.T) {
	party := &partyFacade{&lightController{}, &musicController{}, &videoController{}}

	assert.Equal(t, "light on play music play video", party.partyOn())
}