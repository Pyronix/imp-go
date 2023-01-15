package types

type MultExpression [2]Expression

func Mult(x, y Expression) Expression {
	return (MultExpression)([2]Expression{x, y})
}

func (e MultExpression) Pretty() string {
	var x string
	x = e[0].Pretty()
	x += " * "
	x += e[1].Pretty()

	return x
}

func (e MultExpression) Eval(s *ValueState) Value {
	n1 := e[0].Eval(s)
	if n1.ValueType != ValueInt {
		return UndefinedValue()
	}
	n2 := e[1].Eval(s)
	if n2.ValueType == ValueInt {
		return IntValue(n1.IntValue * n2.IntValue)
	}
	return UndefinedValue()
}

func (e MultExpression) Infer(t *TypeState) Type {
	if t1 := e[0].Infer(t); t1 != TypeInt {
		return TypeIllTyped
	}
	if t2 := e[1].Infer(t); t2 == TypeInt {
		return TypeInt
	}
	return TypeIllTyped
}
