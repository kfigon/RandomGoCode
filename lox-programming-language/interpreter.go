package main

import (
	"fmt"
	"strconv"
)

type loxObject struct {
	v *any
}

type interpreter struct {
	lastResult loxObject
	lastError  error
}

func interpret(expr []expression) ([]loxObject, error) {
	var out []loxObject
	i := &interpreter{}
	for _, e := range expr {
		e.visitExpr(i)
		if i.lastError != nil {
			return nil, i.lastError
		}
		out = append(out, i.lastResult)
	}
	return out, nil
}

func (i *interpreter) visitLiteral(li literal) {
	tok := token(li)
	if checkTokenType(tok, number) {
		v, err := strconv.Atoi(li.lexeme)
		if err != nil {
			i.lastError = fmt.Errorf("invalid number %v, line %v, error: %w", li, li.line, err)
			return 
		}
		i.lastResult = loxObject{v: toAnyPtr(v)}
	} else if checkTokenType(tok, stringLiteral) {
		i.lastResult = loxObject{v: toAnyPtr(li.lexeme)}
	} else if checkTokenType(tok, boolean) {
		v, err := strconv.ParseBool(li.lexeme)
		if err != nil {
			i.lastError = fmt.Errorf("invalid boolean %v, line %v, error: %w", li, li.line, err)
			return 
		}
		i.lastResult = loxObject{v: toAnyPtr(v)}
	} else {
		i.lastError = fmt.Errorf("invalid literal %v, line %v", li, li.line)
	}
}
func (i *interpreter) visitUnary(u unary) {
	op := u.op.lexeme

	tmp := &interpreter{}
	u.ex.visitExpr(tmp)
	if tmp.lastError != nil {
		i.lastError = tmp.lastError
		return
	} else if tmp.lastResult.v == nil {
		i.lastError = fmt.Errorf("failed to evaluate unary expression %v, line %v", u.op, u.op.line)
		return
	}

	if op == "!" {
		b, ok := (*tmp.lastResult.v).(*bool)
		if !ok {
			i.lastError = fmt.Errorf("boolean value not found %v, line %v", u.op, u.op.line)
			return
		}
		i.lastResult = loxObject{v: toAnyPtr(!*b)}
	} else if op == "-" {
		b, ok := (*tmp.lastResult.v).(*int)
		if !ok {
			i.lastError = fmt.Errorf("number value not found %v, line %v", u.op, u.op.line)
			return
		}
		i.lastResult = loxObject{v: toAnyPtr(-(*b))}
	} else {
		i.lastError = fmt.Errorf("invalid unary operator %v, line %v", u.op, u.op.line)
		return
	}
}
func (i *interpreter) visitBinary(b binary) {
	leftI := &interpreter{}
	rightI := &interpreter{}
	b.left.visitExpr(leftI)
	b.right.visitExpr(rightI)

	if leftI.lastError != nil {
		i.lastError = leftI.lastError
		return
	} else if rightI.lastError != nil {
		i.lastError = rightI.lastError
		return
	}  else if leftI.lastResult.v == nil {
		i.lastError = fmt.Errorf("failed to evaluate left binary expression %v, line %v", b.op, b.op.line)
		return
	} else if rightI.lastResult.v == nil {
		i.lastError = fmt.Errorf("failed to evaluate right binary expression %v, line %v", b.op, b.op.line)
		return
	}
	isNumber := func(v *any) (int, bool) {
		val, ok := (*v).(*int)
		if !ok {
			i.lastError = fmt.Errorf("number value not found %v, line %v", b.op, b.op.line)
			return 0, false
		}
		return *val, true
	}

	isBool := func(v *any) (bool, bool) {
		val, ok := (*v).(*bool)
		if !ok {
			i.lastError = fmt.Errorf("boolean value not found %v, line %v", b.op, b.op.line)
			return false, false
		}
		return *val, true
	}

	numeric := map[string]func(int,int)int {
		"+": func(a, b int) int {return a+b},
		"-": func(a, b int) int {return a-b},
		"*": func(a, b int) int {return a*b},
		"/": func(a, b int) int {return a/b},
	}

	boolean := map[string]func(bool,bool)bool {
		"!=": func(a, b bool) bool {return a!=b},
		"==": func(a, b bool) bool {return a==b},
	}

	comparison := map[string]func(int,int)bool {
		">": func(a, b int) bool {return a>b},
		">=": func(a, b int) bool {return a>=b},
		"<": func(a, b int) bool {return a<b},
		"<=": func(a, b int) bool {return a<=b},
	}


	if op, ok := boolean[b.op.lexeme]; ok {
		leftV,leftOk := isBool(leftI.lastResult.v)
		rightV,rightOk := isBool(rightI.lastResult.v)
		if leftOk && rightOk {
			i.lastResult = loxObject{v: toAnyPtr(op(leftV, rightV))}
		}
	} else {
		leftV,leftOk := isNumber(leftI.lastResult.v)
		rightV,rightOk := isNumber(rightI.lastResult.v)
		if !leftOk || !rightOk {
			return
		}

		if op, ok := numeric[b.op.lexeme]; ok {
			i.lastResult = loxObject{v: toAnyPtr(op(leftV, rightV))}
		} else if op,ok := comparison[b.op.lexeme]; ok {
			i.lastResult = loxObject{v: toAnyPtr(op(leftV, rightV))}
		} else {
			i.lastError = fmt.Errorf("unsupported binary operator %v, line %v", b.op, b.op.line)
		}
	}
}

func toAnyPtr[T any](v T) *any {
	out := any(v)
	return &out
}