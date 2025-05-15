package main

import "fmt"

type evaluator struct{}

func Eval(program []Statement) (Object, error) {
	e := &evaluator{}
	return e.processStatements(program)
}

func (e *evaluator) processStatements(stmts []Statement) (Object, error) {
	var lastObj Object
	for _, stm := range stmts {
		var err error
		lastObj, err = e.evalNode(stm)
		if err != nil {
			return nil, err
		}
	}
	return lastObj, nil
}

func (e *evaluator) evalNode(st Statement) (Object, error) {
	switch vs := st.(type) {
	case *LetStatement: todo()
	case *ReturnStatement: return e.evalExp(vs.Exp)
	case *BlockStatement: return e.processStatements(vs.Stmts)
	case *ExpressionStatement: return e.evalExp(vs.Exp)
	}
	return nil, fmt.Errorf("invalid node: %T", st)
}

func (e *evaluator) evalExp(vs Expression) (Object, error) {
	switch exp := vs.(type) {
	case *PrimitiveLiteral[int]: return &PrimitiveObj[int]{exp.Val}, nil
	case *PrimitiveLiteral[bool]: return toBool(exp.Val), nil
	case *IdentifierExpression: 
		if exp.Name == "null" {
			return NULL, nil
		}
		todo()
	case *PrefixExpression: 
		evaluated, err := e.evalExp(exp.Expr)
		if err != nil {
			return nil, err
		} else if exp.Operator.Typ == Bang {
			b, ok := evaluated.(*PrimitiveObj[bool])
			if !ok{
				return nil, fmt.Errorf("expected boolean type, got %T", evaluated)
			}
			return toBool(!b.Data),nil
		} else if exp.Operator.Typ == Minus {
			i, ok := evaluated.(*PrimitiveObj[int])
			if !ok{
				return nil, fmt.Errorf("expected int type, got %T", evaluated)
			}
			return &PrimitiveObj[int]{-i.Data},nil
		}
		return nil, fmt.Errorf("unsupported prefix operator: %v", exp.Operator.Typ)
	case *InfixExpression: 
		left, err := e.evalExp(exp.Left)
		if err != nil {
			return nil, err
		}
		right, err := e.evalExp(exp.Right)
		if err != nil {
			return nil, err
		}

		if iL, iR, ok := castBothToPrimitive[int](left,right); ok{
			switch exp.Operator.Typ{
			case Plus: return &PrimitiveObj[int]{iL+iR}, nil
			case Minus: return &PrimitiveObj[int]{iL-iR}, nil
			case Asterisk: return &PrimitiveObj[int]{iL*iR}, nil
			case Slash: return &PrimitiveObj[int]{iL/iR}, nil
			case LT: return &PrimitiveObj[bool]{iL < iR}, nil
			case GT: return &PrimitiveObj[bool]{iL > iR}, nil
			default: return nil, fmt.Errorf("invalid operator for integers: %v", exp.Operator.Typ)
			}
		} else if bL, bR, ok := castBothToPrimitive[bool](left,right); ok{
			switch exp.Operator.Typ{
			case EQ: return &PrimitiveObj[bool]{bL == bR}, nil
			case NEQ: return &PrimitiveObj[bool]{bL != bR}, nil
			default: return nil, fmt.Errorf("invalid operator for booleans %T, %T", left, right)
			}
		}
		return nil, fmt.Errorf("expected int or boolean expression, got %T, %T", left, right)
	case *IfExpression: todo()
	case *FunctionLiteral: todo()
	case *FunctionCall: todo()
	}

	return nil, fmt.Errorf("invalid expression: %T", vs)
}

func todo() {
	panic("not implemented")
}

func castBothToPrimitive[T any](a,b Object) (T, T, bool) {
	a2, aOk := a.(*PrimitiveObj[T])
	b2, bOk := b.(*PrimitiveObj[T])
	if aOk && bOk {
		return a2.Data,b2.Data, true
	}
	var zero T
	return zero, zero, false
}

func toBool(v bool) Object {
	if v {
		return TRUE
	}
	return FALSE
}