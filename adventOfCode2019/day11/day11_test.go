package day11

import (
	"aoc2019/intcode"
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
	controller := newRobot()

	controller.process(1,0)
	assert.Equal(t, position{-1,0},controller.position)

	controller.process(0,0)
	assert.Equal(t, position{-1,-1},controller.position)

	controller.process(1,0)
	controller.process(1,0)
	assert.Equal(t, position{0,0},controller.position)
	assert.Equal(t, COLOR_WHITE, controller.currentColor())

	controller.process(0,1)
	controller.process(1,0)
	controller.process(1,0)

	assert.Equal(t, position{0,1},controller.position)

	assert.Equal(t, 6, len(controller.grid))
}

func TestPart1(t *testing.T) {
	computer := intcode.NewComputer(readFile(t))
	controller := newRobot()

	twoOutputs := []int{}

	for !computer.SingleInstruction() {

		if computer.NextInput() {
			computer.SetUserInput(controller.currentColor())
		} else if computer.NextOuput() {
			computer.SingleInstruction()
			nextColor := computer.GetOutput()
			twoOutputs = append(twoOutputs, nextColor)
		}

		if len(twoOutputs) == 2 {
			controller.process(twoOutputs[0], twoOutputs[1])
			twoOutputs = []int{}
		}
	}
	result := len(controller.grid)
	assert.NotEqual(t, 1521, result)
	assert.NotEqual(t, 6016, result)
	assert.NotEqual(t, 1248, result)
	assert.NotEqual(t, 8218, result)
	assert.NotEqual(t, 2194, result)
	assert.Equal(t, 6, result)
}


type position struct { x,y int }

type robot struct {
	position
	direction int
	grid map[position]int
}

func newRobot() *robot {
	return &robot{
		grid: map[position]int{},
	}
}


const (
	DIRECTION_UP = iota
	DIRECTION_DOWN
	DIRECTION_LEFT
	DIRECTION_RIGHT
)
const (
	COLOR_BLACK = iota
	COLOR_WHITE
)

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

const (
	MOVE_LEFT = iota
	MOVE_RIGHT
)
func (this *robot) process(paintInstruction, moveInstruction int) {
	this.grid[this.position] = paintInstruction

	if moveInstruction == MOVE_LEFT {
		this.left()
	} else if moveInstruction == MOVE_RIGHT {
		this.right()
	}
}

func (this *robot) currentColor() int {
	return this.grid[this.position]
}