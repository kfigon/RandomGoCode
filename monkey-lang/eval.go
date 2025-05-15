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
		if exp.Val {
			return TRUE,nil
		} else {
			return FALSE,nil
		}
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
			if b.Data {
				return FALSE,nil
			}
			return TRUE, nil
		} else if exp.Operator.Typ == Minus {
			i, ok := evaluated.(*PrimitiveObj[int])
			if !ok{
				return nil, fmt.Errorf("expected int type, got %T", evaluated)
			}
			return &PrimitiveObj[int]{-i.Data},nil
		} else {
			return nil, fmt.Errorf("unsupported prefix operator: %v", exp.Operator.Typ)
		}
	case *InfixExpression: return nil, nil
	case *IfExpression: return nil, nil
	case *FunctionLiteral: return nil, nil
	case *FunctionCall: return nil, nil
	}

	return nil, fmt.Errorf("invalid expression: %T", vs)
}

func mustCast[T any](v any) T {
	return v.(T)
}