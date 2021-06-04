package intcode

type Computer []int
func (c Computer) Calc() []int {
	var i int
	for i < len(c) && i != -1 {
		i = c.handleCommand(i)
	}
	return c
}

const (
	ADD = 1
	MULT = 2
	TERMINATE = 99
)
func (c Computer) handleCommand(idx int) int {
	val := c[idx]
	switch val {
	case ADD:	return c.handleAdd(idx)
	case MULT:	return c.handleMult(idx)
	case TERMINATE: return -1
	}
	// should not happen
	return -1
}

func (c Computer) getIdxs(idx int) (int,int,int) {
	return c[idx+1],c[idx+2],c[idx+3]
}

func (c Computer) handleAdd(idx int) int {
	idxA,idxB,idxC := c.getIdxs(idx)
	c[idxC] = c[idxA] + c[idxB]
	return idx+4
}

func (c Computer) handleMult(idx int) int {
	idxA,idxB,idxC := c.getIdxs(idx)
	c[idxC] = c[idxA] * c[idxB]
	return idx+4
}