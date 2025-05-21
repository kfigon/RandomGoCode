package compiler

import "monkey-lang/parser"

type Compiler struct {
	instructions Instructions
	constants any // todo: define
}

func NewCompiler() *Compiler {
	return &Compiler{
		instructions: Instructions{},
		constants: nil,
	}
}

func Compile(program []parser.Statement) (Instructions, error) {
	c := NewCompiler()
	return c.compile(program)
}

func (c *Compiler) compile(program []parser.Statement) (Instructions, error) {
	return nil, nil
}
