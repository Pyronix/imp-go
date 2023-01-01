package statements

import (
	"fmt"
	. "imp/types"
)

type IfThenElseStatement struct {
	cond          Expression
	thenBlockStmt BlockStatement
	elseBlockStmt BlockStatement
}

func Ite(x Expression, y BlockStatement, z BlockStatement) Statement {
	return IfThenElseStatement{x, y, z}
}

func (ite IfThenElseStatement) Eval(s ValueState) {
	v := ite.cond.Eval(s)
	if v.ValueType == ValueBool {
		switch {
		case v.BoolValue:
			ite.thenBlockStmt.Eval(s)
		case !v.BoolValue:
			ite.elseBlockStmt.Eval(s)
		}
	} else {
		fmt.Printf("if-then-else Eval fail")
	}
}

func (ite IfThenElseStatement) Pretty() string {
	var x string
	x = "if "
	x += ite.cond.Pretty()
	x += " "
	x += ite.thenBlockStmt.Pretty()
	x += " else "
	x += ite.elseBlockStmt.Pretty()

	return x
}

func (ite IfThenElseStatement) Check(t TypeState) bool {
	ty := ite.cond.Infer(t)
	if ty != TypeBool {
		return false
	}

	return ite.thenBlockStmt.Check(t) && ite.elseBlockStmt.Check(t)
}
