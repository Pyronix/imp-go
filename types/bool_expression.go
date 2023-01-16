package types

type BoolExpression bool

func Bool(x bool) Expression {
	return BoolExpression(x)
}

func (x BoolExpression) Pretty() string {
	if x {
		return "true"
	} else {
		return "false"
	}
}

func (x BoolExpression) Eval(s *ValueState) Value {
	return BoolValue((bool)(x))
}

func (x BoolExpression) Infer(t *TypeState) Type {
	return TypeBool
}
