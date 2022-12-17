package types

/*
type NegationExpression Expression

func Negation(x Expression) Expression {
	return (NegationExpression)(x)
}

func (e NegationExpression) Pretty() string {
	var x string
	x = "(!"
	x += e.Pretty()
	x += ")"

	return x
}

func (e NegationExpression) Eval(s ValueState) Value {
	b := e.Eval(s)
	switch {
	case b.ValueType == ValueBool:
		return BoolValue(!b.BoolValue)
	}
	return UndefinedValue()
}

func (e NegationExpression) Infer(t TypeState) Type {
	t1 := e.Infer(t)
	if t1 == TypeBool {
		return TypeBool
	}
	return TypeIllTyped
}
*/
