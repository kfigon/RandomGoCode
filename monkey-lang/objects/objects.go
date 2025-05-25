package objects
import (
	"monkey-lang/parser"
)

type Object interface {
	objTag()
}

var (
	NULL = &NullObj{}
	TRUE = &PrimitiveObj[bool]{true}
	FALSE = &PrimitiveObj[bool]{false}
)

type PrimitiveObj[T any] struct {
	Data T
}

func (*PrimitiveObj[T]) objTag(){}

type NullObj struct{}
func (*NullObj) objTag(){}

type ReturnObj struct {
	Val Object
}

func (*ReturnObj) objTag(){}

type FunctionObj struct {
	Args []string
	Body *parser.BlockStatement
}

func (*FunctionObj) objTag(){}
