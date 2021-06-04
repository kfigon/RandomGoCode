package intcode

import "math"

type Computer struct {
	instructions []int
	userInput    int
	userOutput   int
}

func NewComputer(instructions []int) *Computer {
	return &Computer{
		instructions: instructions,
	}
}

const (
	OP_ADD       = 1
	OP_MULT      = 2
	OP_INPUT     = 3
	OP_OUTPUT    = 4
	OP_TERMINATE = 99

	IDX_TERMINATE = -1

	MODE_POSITION  = 0
	MODE_IMMEDIATE = 1
)

func (c *Computer) Calc() []int {
	var i int
	for i < len(c.instructions) && i != IDX_TERMINATE {
		i = c.handleCommand(i)
	}
	return c.instructions
}

func (c *Computer) handleCommand(idx int) int {
	opcode := opcode(c.instructions[idx])
	switch opcode.extractOpcode() {
	case OP_ADD:
		return c.handleAdd(idx)
	case OP_MULT:
		return c.handleMult(idx)
	case OP_INPUT:
		return c.handleInput(idx)
	case OP_OUTPUT:
		return c.handleOutput(idx)
	case OP_TERMINATE:
		return IDX_TERMINATE
	}
	// should not happen
	return IDX_TERMINATE
}

func (c *Computer) handleAdd(idx int) int {
	op := opcode(c.instructions[idx])
	param0, param1,param2 := c.instructions[idx+1], c.instructions[idx+2], c.instructions[idx+3]
	c.instructions[param2] = c.paramValue(op.modeForParam(0),param0) + c.paramValue(op.modeForParam(1),param1)
	return idx + 4
}

func (c *Computer) handleInput(idx int) int {
	param0 := c.instructions[idx+1]
	c.instructions[param0] = c.userInput
	return idx + 2
}

func (c *Computer) handleOutput(idx int) int {
	op := opcode(c.instructions[idx])
	param0 := c.instructions[idx+1]
	c.userOutput = c.paramValue(op.modeForParam(0),param0)
	return idx + 2
}

func (c *Computer) handleMult(idx int) int {
	op := opcode(c.instructions[idx])
	param0, param1,param2 := c.instructions[idx+1], c.instructions[idx+2], c.instructions[idx+3]
	c.instructions[param2] = c.paramValue(op.modeForParam(0),param0) * c.paramValue(op.modeForParam(1),param1)
	return idx + 4
}

func (c *Computer) paramValue(paramMode int, value int) int {
	if paramMode == MODE_POSITION {
		return c.instructions[value]
	} else if paramMode == MODE_IMMEDIATE {
		return value
	}
	// should never happen
	return -1
}

func (c *Computer) SetUserInput(val int)  {
	c.userInput = val
}

func (c *Computer) GetOutput() int {
	return c.userOutput
}

type opcode int

func (op opcode) extractOpcode() int {
	return int(op) % 100
}

func (op opcode) modeForParam(paramIdx int) int {
	paramFlags := int(op) / 100
	return (paramFlags / int(math.Pow10(paramIdx))) % 10
}