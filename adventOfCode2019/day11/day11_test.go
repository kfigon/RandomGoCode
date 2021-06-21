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
	result := part1(readFile(t))
	
	assert.Equal(t, 2415, result)
}

func part1(input []int) int {
	computer := intcode.NewComputer(input)
	controller := newRobot()
	lastTwoOutputs := []int{}
	
	for !computer.SingleInstruction() {

		if computer.NextInput() {
			computer.ClearUserInput()
			computer.SetUserInput(controller.currentColor())
		} else if computer.NextOuput() {
			computer.SingleInstruction()
			ouputVal := computer.GetOutput()
			lastTwoOutputs = append(lastTwoOutputs, ouputVal)
		}

		if len(lastTwoOutputs) == 2 {
			controller.process(lastTwoOutputs[0], lastTwoOutputs[1])
			lastTwoOutputs = []int{}
		}
	}
	return len(controller.grid)
}


type position struct { x,y int }
type robot struct {
	position
	direction int
	grid map[position]int
}

func newRobot() *robot {
	return &robot{ grid: map[position]int{}}
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
const (
	MOVE_LEFT = iota
	MOVE_RIGHT
)

func (r *robot) left() {
	switch r.direction {
	case DIRECTION_UP: 
		r.direction = DIRECTION_LEFT
		r.x--
	case DIRECTION_DOWN: 
		r.direction = DIRECTION_RIGHT
		r.x++
	case DIRECTION_LEFT:
		r.direction = DIRECTION_DOWN
		r.y--
	case DIRECTION_RIGHT:
		r.direction = DIRECTION_UP
		r.y++
	}
}

func (r *robot) right() {
	switch r.direction {
	case DIRECTION_UP: 
		r.direction = DIRECTION_RIGHT
		r.x++
	case DIRECTION_DOWN: 
		r.direction = DIRECTION_LEFT
		r.x--
	case DIRECTION_LEFT:
		r.direction = DIRECTION_UP
		r.y++
	case DIRECTION_RIGHT:
		r.direction = DIRECTION_DOWN
		r.y--
	}
}

func (r *robot) process(paintInstruction, moveInstruction int) {
	r.grid[r.position] = paintInstruction

	if moveInstruction == MOVE_LEFT {
		r.left()
	} else if moveInstruction == MOVE_RIGHT {
		r.right()
	}
}

func (r *robot) currentColor() int {
	return r.grid[r.position]
}