package types

type OrExpression [2]Expression

func Or(x, y Expression) Expression {
	return (OrExpression)([2]Expression{x, y})
}

func (e OrExpression) Pretty() string {
	var x string

	x = e[0].Pretty()
	x += " || "
	x += e[1].Pretty()

	return x
}

func (e OrExpression) Eval(s ValueState) Value {
	b1 := e[0].Eval(s)
	b2 := e[1].Eval(s)
	switch {
	case b1.ValueType == ValueBool && b1.BoolValue == true:
		return BoolValue(true)
	case b1.ValueType == ValueBool && b2.ValueType == ValueBool:
		return BoolValue(b1.BoolValue || b2.BoolValue)
	}
	return UndefinedValue()
}

func (e OrExpression) Infer(t TypeState) Type {
	t1 := e[0].Infer(t)
	t2 := e[1].Infer(t)
	if t1 == TypeBool && t2 == TypeBool {
		return TypeBool
	}
	return TypeIllTyped
}
