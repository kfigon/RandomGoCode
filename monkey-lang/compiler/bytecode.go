package compiler

import (
	"encoding/binary"
	"fmt"
	"iter"
)

type Instructions []byte

type Opcode byte
const (
	OpConst Opcode = iota
)

func(o Opcode) String() string {
	return fmt.Sprintf("%02x %s", byte(o), opcodeLookup[o].name)
}

type opcodeDefinition struct {
	name string
	operandWidth []int
}

var opcodeLookup = map[Opcode]opcodeDefinition {
	OpConst: {"OpConst", []int{2}},
}

var endianness = binary.BigEndian

func MakeCommand(op Opcode, operands ...int) ([]byte, error) {
	def, ok := opcodeLookup[op] 
	if !ok {
		return nil, fmt.Errorf("unknown opcode %x", op)
	} else if len(operands) != len(def.operandWidth) {
		return nil, fmt.Errorf("not matched number of operands. %d given for op %v -> %d", len(operands), op, len(def.operandWidth))
	}

	instructionLen := 1
	for _, w := range def.operandWidth {
		instructionLen += w
	}

	instr := make([]byte, instructionLen)
	instr[0] = byte(op)

	offset := 1
	for i, operand := range operands {
		width := def.operandWidth[i]
		if width == 1 {
			instr[offset] = byte(operand)
		} else if width == 2 {
			endianness.PutUint16(instr[offset:], uint16(operand))
		}

		offset += width
	}

	return instr, nil
}

func (ins Instructions) Iter() iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for i := 0; i < len(ins); {
			nextOffset := i + 1
			for _, v := range opcodeLookup[Opcode(ins[i])].operandWidth {
				nextOffset += v
			}	

			if !yield(ins[i:nextOffset]) {
				return
			}
			i += nextOffset
		}
	}
}