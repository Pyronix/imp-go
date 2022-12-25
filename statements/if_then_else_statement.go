package statements

import (
	"fmt"
	. "imp/types"
)

type IfThenElseStatement struct {
	cond     Expression
	thenStmt Statement
	elseStmt Statement
}

func (ite IfThenElseStatement) Eval(s ValueState) {
	v := ite.cond.Eval(s)
	if v.ValueType == ValueBool {
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

func (ite IfThenElseStatement) Pretty(s ValueState) string {
	var x string
	x = "if "
	x += ite.cond.Pretty()
	x += " { "
	x += ite.thenStmt.Pretty()
	x += " } else { "
	x += ite.elseStmt.Pretty()
	x += " }"

	return x
}

func (ite IfThenElseStatement) Check(t TypeState) bool {
	ty := ite.cond.Infer(t)
	if ty != TypeBool {
		return false
	}

	return ite.thenStmt.Check(t) && ite.elseStmt.Check(t)
}
