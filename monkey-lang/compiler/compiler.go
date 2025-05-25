package compiler

import (
	"fmt"
	"monkey-lang/objects"
	"monkey-lang/parser"
)

type Compiler struct {
	instructions Instructions
	constants []objects.Object
}

func NewCompiler() *Compiler {
	return &Compiler{
		instructions: Instructions{},
		constants: nil,
	}
}

func Compile(program []parser.Statement) (Instructions, []objects.Object, error){
	c := NewCompiler()
	err := c.compile(program)
	if err != nil {
		return nil, nil, err
	}
	return c.instructions, c.constants, nil
}

func (c *Compiler) compile(program []parser.Statement) error {
	for _, st := range program {
		switch v := st.(type) {
		case *parser.ExpressionStatement: return c.compileExpression(v.Exp)
		default: return fmt.Errorf("invalid type %T", st)
		}
	}
	return nil
}

func (c *Compiler) compileExpression(exp parser.Expression) error {
	switch v := exp.(type) {
	case *parser.PrimitiveLiteral[int]:
		instr, err := MakeCommand(OpConst, len(c.constants))
		if err != nil {
			return err
		}
		c.constants = append(c.constants, &objects.PrimitiveObj[int]{Data: v.Val})
		c.instructions = append(c.instructions, instr...)
	default: return fmt.Errorf("invalid expression type %T", exp)
	}
	return nil
}
