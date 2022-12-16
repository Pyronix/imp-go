package types

import (
	"strconv"
)

type NumberExpression int

func Number(x int) Expression {
	return NumberExpression(x)
}

func (x NumberExpression) Pretty() string {
	return strconv.Itoa(int(x))
}

func (x NumberExpression) Eval(s ValueState) Value {
	return IntValue((int)(x))
}

func (x NumberExpression) Infer(t TypeState) Type {
	return TypeInt
}
