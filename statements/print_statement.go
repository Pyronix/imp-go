package statements

import . "imp/types"

type PrintStatement struct {
	Expression
}

func Print(x Expression) Statement {
	return PrintStatement{x}
}

func (print PrintStatement) Pretty() string {
	return "print " + print.Expression.Pretty()
}

func (print PrintStatement) Eval(s *ValueState) {
	print.Expression.Eval(s)
}

func (print PrintStatement) Check(t *TypeState) bool {
	return print.Expression.Infer(t) != TypeIllTyped
}
