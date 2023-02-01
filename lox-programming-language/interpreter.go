package main

import (
	"fmt"
	"strconv"
)

type loxObject struct {
	v *any
}

type interpreter struct {
	// go is stupid with generics, so this is a workaround to get results
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
		} else {
			i.lastResult = loxObject{v: toAnyPtr(v)}
		}
	} else if checkTokenType(tok, stringLiteral) {
		i.lastResult = loxObject{v: toAnyPtr(li.lexeme)}
	} else if checkTokenType(tok, boolean) {
		v, err := strconv.ParseBool(li.lexeme)
		if err != nil {
			i.lastError = fmt.Errorf("invalid boolean %v, line %v, error: %w", li, li.line, err)
		} else {
			i.lastResult = loxObject{v: toAnyPtr(v)}
		}
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
		v, err := castTo[bool](u.op, tmp.lastResult.v)
		if err != nil {
			i.lastError = err
		} else {
			i.lastResult = loxObject{v: toAnyPtr(!v)}
		}
	} else if op == "-" {
		v, err := castTo[int](u.op, tmp.lastResult.v)
		if err != nil {
			i.lastError = err
		} else {
			i.lastResult = loxObject{v: toAnyPtr(-v)}
		}
	} else {
		i.lastError = fmt.Errorf("invalid unary operator %v, line %v", u.op, u.op.line)
	}
}

func castTo[T any](t token, v *any) (T, error) {
	val, ok := (*v).(T)
	if !ok {
		return val, fmt.Errorf("%v value not found %v, line %v", t.tokType, t, t.line)
	}
	return val, nil
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
		"!=": func(a, b int) bool {return a!=b},
		"==": func(a, b int) bool {return a==b},
	}

	leftBool,leftErr := castTo[bool](b.op, leftI.lastResult.v)
	rightBool,rightErr := castTo[bool](b.op, rightI.lastResult.v)

	if op, ok := boolean[b.op.lexeme]; ok && leftErr == nil && rightErr == nil {
		if leftErr != nil {
			i.lastError = leftErr
		} else if rightErr != nil {
			i.lastError = rightErr
		} else {
			i.lastResult = loxObject{v: toAnyPtr(op(leftBool, rightBool))}
		}
		return
	}

	leftStr, leftErr := castTo[string](b.op, leftI.lastResult.v)
	rightStr, rightErr := castTo[string](b.op, rightI.lastResult.v)
	if leftErr == nil && rightErr == nil {
		if b.op.lexeme == "+" {
			i.lastResult = loxObject{v: toAnyPtr(leftStr + rightStr)}
		} else if b.op.lexeme == "==" {
			i.lastResult = loxObject{v: toAnyPtr(leftStr == rightStr)}
		} else if b.op.lexeme == "!=" {
			i.lastResult = loxObject{v: toAnyPtr(leftStr != rightStr)}
		} else {
			i.lastError = fmt.Errorf("unsupported binary operator on strings %v, line %v", b.op, b.op.line)
		}
		return
	}

	leftV,leftErr := castTo[int](b.op, leftI.lastResult.v)
	rightV,rightErr := castTo[int](b.op, rightI.lastResult.v)
	if leftErr != nil {
		i.lastError = leftErr
	} else if rightErr != nil {
		i.lastError = rightErr
	} else if op, ok := numeric[b.op.lexeme]; ok {
		i.lastResult = loxObject{v: toAnyPtr(op(leftV, rightV))}
	} else if op,ok := comparison[b.op.lexeme]; ok {
		i.lastResult = loxObject{v: toAnyPtr(op(leftV, rightV))}
	} else {
		i.lastError = fmt.Errorf("unsupported binary operator %v, line %v", b.op, b.op.line)
	}
}

func toAnyPtr[T any](v T) *any {
	out := any(v)
	return &out
}