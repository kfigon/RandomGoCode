package day11

import (
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/11
func readFile(t *testing.T) []int {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)

	splitted := strings.Split(string(content),",")
	out := make([]int, 0)
	for i := 0; i < len(splitted); i++ {
		v, err := strconv.Atoi(string(splitted[i]))
		require.NoError(t, err)
		
		out = append(out, v)
	}
	return out	
}

func TestRobot(t *testing.T) {
	controller := newController()

	controller.process(1,0)
	controller.process(0,0)
	controller.process(1,0)
	controller.process(1,0)
	
	controller.process(0,1)
	controller.process(1,0)
	controller.process(1,0)

	assert.Equal(t, 6, len(controller.grid.m))
}

const (
	DIRECTION_UP = iota
	DIRECTION_DOWN
	DIRECTION_LEFT
	DIRECTION_RIGHT
)

type position struct { x,y int }

const (
	COLOR_BLACK = iota
	COLOR_WHITE
)
type grid struct {
	m map[position]int
}

func newGrid() *grid {
	return &grid{
		m: map[position]int{},
	}
}

type robot struct {
	position
	direction int
}

func newRobot() *robot {
	return &robot{}
}

func (this *robot) left() {
	switch this.direction {
	case DIRECTION_UP: 
		this.direction = DIRECTION_LEFT
		this.x--
	case DIRECTION_DOWN: 
		this.direction = DIRECTION_RIGHT
		this.x++
	case DIRECTION_LEFT:
		this.direction = DIRECTION_DOWN
		this.y--
	case DIRECTION_RIGHT:
		this.direction = DIRECTION_UP
		this.y++
	}
}

func (this *robot) right() {
	switch this.direction {
	case DIRECTION_UP: 
		this.direction = DIRECTION_RIGHT
		this.x++
	case DIRECTION_DOWN: 
		this.direction = DIRECTION_LEFT
		this.x--
	case DIRECTION_LEFT:
		this.direction = DIRECTION_UP
		this.y++
	case DIRECTION_RIGHT:
		this.direction = DIRECTION_DOWN
		this.y--
	}
}

type robotController struct {
	robot *robot
	grid *grid
}

func newController() *robotController {
	return &robotController{
		robot: newRobot(),
		grid: newGrid(),
	}
}

const (
	MOVE_LEFT = iota
	MOVE_RIGHT
)
func (this *robotController) process(paintInstruction, moveInstruction int) {
	this.grid.m[this.robot.position] = paintInstruction

	if moveInstruction == MOVE_LEFT {
		this.robot.left()
	} else if moveInstruction == MOVE_RIGHT {
		this.robot.right()
	}
}

func (this *robotController) currentColor() int {
	return this.grid.m[this.robot.position]
}