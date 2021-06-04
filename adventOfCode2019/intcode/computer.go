package intcode

type Computer struct {
	instructions []int
	userInput *int
	userOutput *int
}

func NewComputer(instructions []int) *Computer {
	return &Computer{
		instructions: instructions,
	}
}

const (
	OP_ADD = 1
	OP_MULT = 2
	OP_INPUT = 3
	OP_OUTPUT = 4
	OP_TERMINATE = 99
	
	IDX_TERMINATE = -1
	
	MODE_POSITION = 0
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
	opcode := c.extractOpcode(c.instructions[idx])
	switch opcode {
	case OP_ADD:	return c.handleAdd(idx)
	case OP_MULT:	return c.handleMult(idx)
	case OP_TERMINATE: return IDX_TERMINATE
	}
	// should not happen
	return IDX_TERMINATE
}

func (c *Computer) getIdxs(idx int) (int,int,int) {
	return c.instructions[idx+1],c.instructions[idx+2],c.instructions[idx+3]
}

func (c *Computer) handleAdd(idx int) int {
	idxA,idxB,idxC := c.getIdxs(idx)
	c.instructions[idxC] = c.instructions[idxA] + c.instructions[idxB]
	return idx+4
}

func (c *Computer) handleMult(idx int) int {
	idxA,idxB,idxC := c.getIdxs(idx)
	c.instructions[idxC] = c.instructions[idxA] * c.instructions[idxB]
	return idx+4
}

func (c *Computer) extractOpcode(val int) int {
	return val % 100
}