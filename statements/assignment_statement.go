package statements

import (
	"fmt"
	"imp/types"
)

type AssignmentStatement struct {
	lhs string
	rhs types.Expression
}

func (ass AssignmentStatement) Eval(s types.ValueState) {
	v := ass.rhs.Eval(s)
	x := (string)(ass.lhs)
	if v.ValueType == s[x].ValueType {
		s[x] = v
	} else {
		fmt.Printf("Assignment Eval fail")
	}
}

func (ass AssignmentStatement) Pretty() string {
	return ass.lhs + " = " + ass.rhs.Pretty()
}

func (ass AssignmentStatement) Check(t types.TypeState) bool {
	ty := ass.rhs.Infer(t)
	if ty == types.TypeIllTyped {
		return false
	}

	x := (string)(ass.lhs)
	if t[x] != ass.rhs.Infer(t) {
		return false
	}
	t[x] = ty
	return true
}
