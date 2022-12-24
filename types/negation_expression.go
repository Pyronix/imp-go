package types

type NegationExpression struct {
	Expression
}

func Negation(x Expression) Expression {
	return NegationExpression{x}
}

func (e NegationExpression) Pretty() string {
	var x string

	x = "!"
	x += e.Expression.Pretty()

	return x
}

func (e NegationExpression) Eval(s ValueState) Value {
	b := e.Expression.Eval(s)
	switch {
	case b.ValueType == ValueBool:
		return BoolValue(!b.BoolValue)
	}
	return UndefinedValue()
}

func (e NegationExpression) Infer(t TypeState) Type {
	t1 := e.Expression.Infer(t)
	if t1 == TypeBool {
		return TypeBool
	}
	return TypeIllTyped
}
