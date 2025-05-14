package main

type Object interface {
	objTag()
}

type PrimitiveObj[T any] struct {
	Data T
}

func (*PrimitiveObj[T]) objTag(){}

type NullObj struct{}
func (*NullObj) objTag(){}


