package statements

import (
	"fmt"
	. "imp/types"
)

type AssignmentStatement struct {
	lhs string
	rhs Expression
}

func (ass AssignmentStatement) Eval(s ValueState) {
	v := ass.rhs.Eval(s)
	x := (string)(ass.lhs)
	val, ok := s[x]
	if ok && v.ValueType == val.ValueType && v.ValueType != Undefined {
		s[x] = v
	} else {
		fmt.Printf("Assignment Eval fail")
	}
}

func (ass AssignmentStatement) Pretty() string {
	return ass.lhs + " = " + ass.rhs.Pretty()
}

func (ass AssignmentStatement) Check(t TypeState) bool {
	ty := ass.rhs.Infer(t)
	if ty == TypeIllTyped {
		return false
	}

	x := (string)(ass.lhs)
	if t[x] != ty {
		return false
	}

	return true
}
