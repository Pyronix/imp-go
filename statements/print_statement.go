package statements

import (
	"fmt"
	. "imp/types"
)

type PrintStatement struct {
	Expression
}

func Print(x Expression) Statement {
	return PrintStatement{x}
}

func (print PrintStatement) Pretty() string {
	return "print " + print.Expression.Pretty()
}

func (print PrintStatement) Eval(s *ValueState) Value {
	value := print.Expression.Eval(s)
	fmt.Println(ShowVal(value))

	return value
}

func (print PrintStatement) Check(t *TypeState) bool {
	return print.Expression.Infer(t) != TypeIllTyped
}
