package statements

import (
	"fmt"
	"imp/types"
)

type IfThenElseStatement struct {
	cond     types.Expression
	thenStmt Statement
	elseStmt Statement
}

func (ite IfThenElseStatement) Eval(s types.ValueState) {
	v := ite.cond.Eval(s)
	if v.ValueType == types.ValueBool {
		switch {
		case v.BoolValue:
			ite.thenStmt.Eval(s)
		case !v.BoolValue:
			ite.elseStmt.Eval(s)
		}

	} else {
		fmt.Printf("if-then-else Eval fail")
	}
}
