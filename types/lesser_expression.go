package types

type LesserExpression [2]Expression

func Lesser(x, y Expression) Expression {
	return (LesserExpression)([2]Expression{x, y})
}

func (e LesserExpression) Pretty() string {
	var x string

	x = e[0].Pretty()
	x += " < "
	x += e[1].Pretty()

	return x
}

func (e LesserExpression) Eval(s *ValueState) Value {
	b1 := e[0].Eval(s)
	if b1.ValueType != ValueInt {
		return UndefinedValue()
	}
	b2 := e[1].Eval(s)
	if b2.ValueType == ValueInt {
		return BoolValue(b1.IntValue < b2.IntValue)
	}
	return UndefinedValue()
}

func (e LesserExpression) Infer(t *TypeState) Type {
	if t1 := e[0].Infer(t); t1 != TypeInt {
		return TypeIllTyped
	}
	if t2 := e[1].Infer(t); t2 == TypeInt {
		return TypeBool
	}
	return TypeIllTyped
}
