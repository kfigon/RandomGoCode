package main

import "fmt"

type evaluator struct{}

type environment struct {
	vals map[string]Object
	outer *environment
}

func (e *environment) set(name string, obj Object) {
	e.vals[name] = obj
}

func (e *environment) get(name string) (Object, bool) {
	if v, ok := e.vals[name]; ok {
		return v, true
	} else if e.outer != nil {
		return e.outer.get(name)
	}
	return NULL, false
}

func Eval(program []Statement) (Object, error) {
	e := &evaluator{}
	env := &environment{
		vals: map[string]Object{},
	}
	return e.processProgram(program, env)
}

func (e *evaluator) processProgram(stmts []Statement, env *environment) (Object, error) {
	var lastObj Object
	for _, stm := range stmts {
		var err error
		lastObj, err = e.evalNode(stm, env)
		if err != nil {
			return nil, err
		} else if ret, ok := lastObj.(*ReturnObj); ok {
			return ret.Val, nil
		}
	}
	return lastObj, nil
}

func (e *evaluator) evalBlockStatement(block *BlockStatement, env *environment) (Object, error) {
	var lastObj Object
	for _, stm := range block.Stmts {
		var err error
		lastObj, err = e.evalNode(stm, env)
		if err != nil {
			return nil, err
		} else if ret, ok := lastObj.(*ReturnObj); ok {
			return ret, nil
		}
	}
	return lastObj, nil
}
func (e *evaluator) evalNode(st Statement, env *environment) (Object, error) {
	switch vs := st.(type) {
	case *LetStatement: 
		ex, err := e.evalExp(vs.Value, env)
		if err != nil {
			return nil, err
		}
		env.set(vs.Ident.Name, ex)
		return NULL, nil
	case *ReturnStatement:
		exp, err := e.evalExp(vs.Exp, env)
		if err != nil {
			return nil, err
		}
		return &ReturnObj{exp},nil
	case *BlockStatement: return e.evalBlockStatement(vs, env)
	case *ExpressionStatement: return e.evalExp(vs.Exp, env)
	}
	return nil, fmt.Errorf("invalid node: %T", st)
}

func (e *evaluator) evalExp(vs Expression, env *environment) (Object, error) {
	switch exp := vs.(type) {
	case *PrimitiveLiteral[int]: return &PrimitiveObj[int]{exp.Val}, nil
	case *PrimitiveLiteral[bool]: return toBool(exp.Val), nil
	case *IdentifierExpression: 
		if exp.Name == "null" {
			return NULL, nil
		}
		v, _ := env.get(exp.Name)
		return v, nil
	case *PrefixExpression: 
		evaluated, err := e.evalExp(exp.Expr, env)
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
		left, err := e.evalExp(exp.Left, env)
		if err != nil {
			return nil, err
		}
		right, err := e.evalExp(exp.Right, env)
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
			case EQ: return &PrimitiveObj[bool]{iL == iR}, nil
			case NEQ: return &PrimitiveObj[bool]{iL != iR}, nil
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
	case *IfExpression: return e.evalIf(exp, env)
	case *FunctionLiteral: todo()
	case *FunctionCall: todo()
	}

	return nil, fmt.Errorf("invalid expression: %T", vs)
}

func (e *evaluator) evalIf(ex *IfExpression, env *environment) (Object, error) {
	if ex == nil {
		return NULL, nil
	}

	pred, err := e.evalExp(ex.Predicate, env)
	if err != nil {
		return nil, err
	}
	b, ok := pred.(*PrimitiveObj[bool])
	if !ok {
		return nil, fmt.Errorf("if statement requires predicate")
	}

	if b.Data {
		return e.evalBlockStatement(ex.Consequence, env)
	}
	if ex.Alternative != nil && ex.Alternative.Predicate == nil {
		return e.evalBlockStatement(ex.Alternative.Consequence, env)
	}
	return e.evalIf(ex.Alternative, env)
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