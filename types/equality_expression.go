package types

type EqualityExpression [2]Expression

func Equal(x, y Expression) Expression {
	return (EqualityExpression)([2]Expression{x, y})
}

func (e EqualityExpression) Pretty() string {
	var x string

	x = e[0].Pretty()
	x += "=="
	x += e[1].Pretty()

	return x
}

func (e EqualityExpression) Eval(s ValueState) Value {
	b1 := e[0].Eval(s)
	b2 := e[1].Eval(s)
	switch {
	case b1.ValueType == ValueBool && b2.ValueType == ValueBool:
		return BoolValue(b1.BoolValue && b2.BoolValue)
	case b1.ValueType == ValueInt && b2.ValueType == ValueInt:
		return BoolValue(b1.IntValue == b2.IntValue)
	}
	return UndefinedValue()
}

func (e EqualityExpression) Infer(t TypeState) Type {
	t1 := e[0].Infer(t)
	t2 := e[1].Infer(t)
	if t1 == TypeBool && t2 == TypeBool {
		return TypeBool
	}
	if t1 == TypeInt && t2 == TypeInt {
		return TypeInt
	}
	return TypeIllTyped
}
