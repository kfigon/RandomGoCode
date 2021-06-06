package intcode

import "math"

type Computer struct {
	instructions []int
	userInput    *inputHandler
	userOutput   int
	instructionCounter int
}

func NewComputer(instructions []int) *Computer {
	return &Computer{
		instructions: instructions,
		userInput: newInputHandler(),
	}
}

const (
	OP_ADD       = 1
	OP_MULT      = 2
	OP_INPUT     = 3
	OP_OUTPUT    = 4
	OP_JMP_T = 5
	OP_JMP_F = 6
	OP_LT = 7
	OP_EQ = 8
	OP_TERMINATE = 99

	IDX_TERMINATE = -1

	MODE_POSITION  = 0
	MODE_IMMEDIATE = 1
)

func (c *Computer) Calc() []int {
	for c.instructionCounter < len(c.instructions) && c.instructionCounter != IDX_TERMINATE {
		c.instructionCounter = c.handleCommand(c.instructionCounter)
	}
	return c.instructions
}

func (c *Computer) CalcTilOutput() bool {

	for c.instructionCounter < len(c.instructions) && c.instructionCounter != IDX_TERMINATE {
		c.instructionCounter = c.handleCommand(c.instructionCounter)
		
		if c.instructionCounter < len(c.instructions) && c.instructionCounter != IDX_TERMINATE &&
			opcode(c.instructions[c.instructionCounter]).extractOpcode() == OP_INPUT {
			break
		}
	}
	return c.instructionCounter == IDX_TERMINATE
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
	case OP_JMP_T:
		return c.handleJumpTrue(idx)
	case OP_JMP_F:
		return c.handleJumpFalse(idx)
	case OP_LT:
		return c.handleLessThan(idx)
	case OP_EQ:
		return c.handleEquals(idx)
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
	c.instructions[param0] = c.userInput.next()
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
func (c *Computer) handleJumpTrue(idx int) int {
	op := opcode(c.instructions[idx])
	param0, param1 := c.instructions[idx+1], c.instructions[idx+2]
	if c.paramValue(op.modeForParam(0), param0) != 0 {
		return c.paramValue(op.modeForParam(1), param1)
	}
	return idx + 3
}
func (c *Computer) handleJumpFalse(idx int) int {
	op := opcode(c.instructions[idx])
	param0, param1 := c.instructions[idx+1], c.instructions[idx+2]
	if c.paramValue(op.modeForParam(0), param0) == 0 {
		return c.paramValue(op.modeForParam(1), param1)
	}
	return idx + 3
}
func (c *Computer) handleLessThan(idx int) int {
	op := opcode(c.instructions[idx])
	param0, param1,param2 := c.instructions[idx+1], c.instructions[idx+2],c.instructions[idx+3]
	v0 := c.paramValue(op.modeForParam(0), param0)
	v1 := c.paramValue(op.modeForParam(1), param1)
	if v0 < v1 {
		c.instructions[param2] = 1
	} else {
		c.instructions[param2] = 0
	}
	return idx + 4
}
func (c *Computer) handleEquals(idx int) int {
	op := opcode(c.instructions[idx])
	param0, param1,param2 := c.instructions[idx+1], c.instructions[idx+2],c.instructions[idx+3]
	v0 := c.paramValue(op.modeForParam(0), param0)
	v1 := c.paramValue(op.modeForParam(1), param1)
	if v0 == v1 {
		c.instructions[param2] = 1
	} else {
		c.instructions[param2] = 0
	}
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
	c.userInput.add(val)
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

type inputHandler struct {
	inputs []int
	counter int
}

func newInputHandler() *inputHandler {
	return &inputHandler{
		inputs: make([]int, 0),
		counter: 0,
	}
}

func (in *inputHandler) next() int {
	if len(in.inputs) == 0 {
		return 0
	}
	out := in.inputs[in.counter]
	in.counter = (in.counter+1) % len(in.inputs)
	return out
}

func (in *inputHandler) add(val int) {
	in.inputs = append(in.inputs, val)
}