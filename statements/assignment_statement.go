package statements

import "imp/types"

type AssignmentStatement struct {
	lhs string
	rhs types.Expression
}

func (a AssignmentStatement) Check(t types.TypeState) bool {
	x := (string)(a.lhs)
	return t[x] == a.rhs.Infer(t)
}
