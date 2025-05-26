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

func CastBothToPrimitive[T any](a,b Object) (T, T, bool) {
	a2, aOk := a.(*PrimitiveObj[T])
	b2, bOk := b.(*PrimitiveObj[T])
	if aOk && bOk {
		return a2.Data,b2.Data, true
	}
	var zero T
	return zero, zero, false
}

func ToBool(v bool) Object {
	if v {
		return TRUE
	}
	return FALSE
}