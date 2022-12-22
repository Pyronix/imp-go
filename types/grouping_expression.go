package types

type GroupingExpression struct {
	Expression
}

func Grouping(x Expression) Expression {
	return GroupingExpression{x}
}

func (e GroupingExpression) Pretty() string {
	var x string
	x = "("
	x += e.Pretty()
	x = ")"

	return x
}

func (e GroupingExpression) Eval(s ValueState) Value {
	b := e.Eval(s)
	switch {
	case b.ValueType == ValueBool:
		return BoolValue(b.BoolValue)
	case b.ValueType == ValueInt:
		return IntValue(b.IntValue)
	}
	return UndefinedValue()
}

func (e GroupingExpression) Infer(t TypeState) Type {
	return e.Infer(t)
}
