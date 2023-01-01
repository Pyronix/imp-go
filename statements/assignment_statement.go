package statements

import (
	. "imp/types"
)

type AssignmentStatement struct {
	lhs string
	rhs Expression
}

func Assignment(x string, y Expression) Statement {
	return AssignmentStatement{x, y}
}

func (ass AssignmentStatement) Eval(s ValueState) {
	s.Assign(ass.lhs, ass.rhs.Eval(s))
}

func (ass AssignmentStatement) Pretty() string {
	return ass.lhs + " = " + ass.rhs.Pretty()
}

func (ass AssignmentStatement) Check(t TypeState) bool {
	ty := ass.rhs.Infer(t)
	if ty == TypeIllTyped {
		return false
	}
	typ, _ := t.LookUpTypeByVariableName(ass.lhs)

	if typ != ty {
		return false
	}

	t.Assign(ass.lhs, ty)
	return true
}
