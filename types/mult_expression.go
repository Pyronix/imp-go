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

// TODO: Müssen beide Expressions evaluiert werden? Auch an anderen Stellen prüfen.
func (e MultExpression) Eval(s *ValueState) Value {
	n1 := e[0].Eval(s)
	n2 := e[1].Eval(s)
	if n1.ValueType == ValueInt && n2.ValueType == ValueInt {
		return IntValue(n1.IntValue * n2.IntValue)
	}
	return UndefinedValue()
}

func (e MultExpression) Infer(t *TypeState) Type {
	t1 := e[0].Infer(t)
	t2 := e[1].Infer(t)
	if t1 == TypeInt && t2 == TypeInt {
		return TypeInt
	}
	return TypeIllTyped
}
