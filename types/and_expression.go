package types

type AndExpression [2]Expression

func And(x, y Expression) Expression {
	return (AndExpression)([2]Expression{x, y})
}

func (e AndExpression) Pretty() string {

	var x string
	x = "("
	x += e[0].Pretty()
	x += "&&"
	x += e[1].Pretty()
	x += ")"

	return x
}

func (e AndExpression) Eval(s ValueState) Value {
	b1 := e[0].Eval(s)
	b2 := e[1].Eval(s)
	switch {
	case b1.ValueType == ValueBool && b1.BoolValue == false:
		return BoolValue(false)
	case b1.ValueType == ValueBool && b2.ValueType == ValueBool:
		return BoolValue(b1.BoolValue && b2.BoolValue)
	}
	return UndefinedValue()
}

func (e AndExpression) Infer(t TypeState) Type {
	t1 := e[0].Infer(t)
	t2 := e[1].Infer(t)
	if t1 == TypeBool && t2 == TypeBool {
		return TypeBool
	}
	return TypeIllTyped
}
