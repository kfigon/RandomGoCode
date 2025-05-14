package main

func Eval(program []Statement) (Object, error) {
	var lastObj Object
	for _, stm := range program {
		o, err := evalNode(stm)
		if err != nil {
			return nil, err
		}
		lastObj = o
	}
	return lastObj, nil
}

func evalNode(st Statement) (Object, error) {
	switch vs := st.(type) {
	case *LetStatement: return nil, nil
	case *ReturnStatement: return nil, nil
	case *BlockStatement: return nil, nil
	case *ExpressionStatement: 
		switch exp := vs.Exp.(type) {
		case *PrimitiveLiteral[int]: return nil, nil
		case *PrimitiveLiteral[bool]: return nil, nil
		case *PrefixExpression: return nil, nil
		case *InfixExpression: return nil, nil
		case *IfExpression: return nil, nil
		case *FunctionLiteral: return nil, nil
		case *FunctionCall: return nil, nil
		default:
			_ = exp
		}
	}
	return &NullObj{}, nil
}