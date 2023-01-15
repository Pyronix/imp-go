package types

type PlusExpression [2]Expression

func Plus(x, y Expression) Expression {
	return (PlusExpression)([2]Expression{x, y})

	// The type PlusExpression is defined as the two element array consisting of Expression elements.
	// PlusExpression and [2]Expression are isomorphic but different types.
	// We first build the AST value [2]Expression{x,y}.
	// Then cast this value (of type [2]Expression) into a value of type PlusExpression.

}

func (e PlusExpression) Pretty() string {
	var x string

	x = e[0].Pretty()
	x += " + "
	x += e[1].Pretty()

	return x
}

func (e PlusExpression) Eval(s *ValueState) Value {
	n1 := e[0].Eval(s)
	if n1.ValueType != ValueInt {
		return UndefinedValue()
	}
	if n2 := e[1].Eval(s); n2.ValueType != ValueInt {
		return UndefinedValue()
	} else {
		return IntValue(n1.IntValue + n2.IntValue)
	}
}

func (e PlusExpression) Infer(t *TypeState) Type {
	if t1 := e[0].Infer(t); t1 != TypeInt {
		return TypeIllTyped
	}
	if t2 := e[1].Infer(t); t2 == TypeInt {
		return TypeInt
	}
	return TypeIllTyped
}
