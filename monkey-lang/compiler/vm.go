package compiler

import (
	"fmt"
	"monkey-lang/objects"
)

type Stack[T any] struct {
	s []T
	pointer int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		s: make([]T, 0, 512),
		pointer: 0,
	}
}

func (s *Stack[T]) Push(v T) {
	s.s = append(s.s, v)
	s.pointer++
}

func (s *Stack[T]) Pop() T {
	s.pointer--
	out := s.s[s.pointer]
	return out
}

func (s *Stack[T]) Empty() bool {
	return s.pointer <= 0
}

type VM struct {
	instructions Instructions
	constants []objects.Object

	stack *Stack[objects.Object]
}

func NewVM(instr Instructions, consts []objects.Object) *VM {
	return &VM{ instr, consts, NewStack[objects.Object]() }
}

func (v *VM) Execute() error {
	for i := range v.instructions.Iter() {
		op := Opcode(i[0]) 
		switch op {
		case OpConst: 
			c := v.constants[int(endianness.Uint16(i[1:]))]
			v.stack.Push(c)
		case OpAdd:
			right := v.stack.Pop()
			left := v.stack.Pop()
			
			a,b, ok := objects.CastBothToPrimitive[int](right, left)
			if !ok {
				return fmt.Errorf("invalid values provided to summary %T, %T", left, right)
			}
			v.stack.Push(&objects.PrimitiveObj[int]{a+b})
		default: return fmt.Errorf("unknown opcode %v", op)
		}
	}
	return nil
}