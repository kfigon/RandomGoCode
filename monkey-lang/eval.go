package main

import "fmt"

type evaluator struct{}

func Eval(program []Statement) (Object, error) {
	e := &evaluator{}
	var lastObj Object
	for _, stm := range program {
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
	case *LetStatement: return nil, nil
	case *ReturnStatement: return e.evalExp(vs.Exp)
	case *BlockStatement: return nil, nil
	case *ExpressionStatement: return e.evalExp(vs.Exp)
	}
	return nil, fmt.Errorf("invalid node: %T", st)
}

func (e *evaluator) evalExp(vs Expression) (Object, error) {
	switch exp := vs.(type) {
	case *PrimitiveLiteral[int]: return &PrimitiveObj[int]{exp.Val}, nil
	case *PrimitiveLiteral[bool]: 
		return toBool(exp.Val), nil
	case *IdentifierExpression: 
		if exp.Name == "null" {
			return NULL, nil
		}
		return nil, nil
	case *PrefixExpression: 
		evaluated, err := e.evalExp(exp.Expr)
		if err != nil {
			return nil, err
		}
		if exp.Operator.Typ == Bang {
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
		} else {
			return nil, fmt.Errorf("unsupported prefix operator: %v", exp.Operator.Typ)
		}
	case *InfixExpression: 
		left, err := e.evalExp(exp.Left)
		if err != nil {
			return nil, err
		}
		right, err := e.evalExp(exp.Right)
		if err != nil {
			return nil, err
		}
		iLeft, ilOk := left.(*PrimitiveObj[int])
		iRight, irOk := right.(*PrimitiveObj[int])
		bLeft, blOk := left.(*PrimitiveObj[bool])
		bRight, brOk := right.(*PrimitiveObj[bool])

		if ilOk && irOk {
			switch exp.Operator.Typ{
			case Plus: return &PrimitiveObj[int]{iLeft.Data+iRight.Data},nil
			case Minus: return &PrimitiveObj[int]{iLeft.Data-iRight.Data},nil
			case Asterisk: return &PrimitiveObj[int]{iLeft.Data*iRight.Data},nil
			case Slash: return &PrimitiveObj[int]{iLeft.Data/iRight.Data},nil
			case LT: return &PrimitiveObj[bool]{iLeft.Data < iRight.Data}, nil
			case GT: return &PrimitiveObj[bool]{iLeft.Data > iRight.Data}, nil
			default: return nil, fmt.Errorf("invalid operator for integers: %v", exp.Operator.Typ)
			}
		} else if blOk && brOk {
			switch exp.Operator.Typ{
			case EQ: return &PrimitiveObj[bool]{bLeft.Data == bRight.Data},nil
			case NEQ: return &PrimitiveObj[bool]{bLeft.Data != bRight.Data},nil
			default: return nil, fmt.Errorf("invalid operator for booleans %T, %T", left, right)
			}
		} else {
			return nil, fmt.Errorf("expected int or boolean expression, got %T, %T", left, right)
		}
	case *IfExpression: return nil, nil
	case *FunctionLiteral: return nil, nil
	case *FunctionCall: return nil, nil
	}

	return nil, fmt.Errorf("invalid expression: %T", vs)
}

func toBool(v bool) Object {
	if v {
		return TRUE
	}
	return FALSE
}