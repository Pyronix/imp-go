package statements

import (
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

func (ite IfThenElseStatement) Eval(s *ValueState) Value {
	v := ite.cond.Eval(s)
	if v.ValueType != ValueBool {
		panic("if-then-else Eval fail")
	}

	if v.BoolValue {
		return ite.thenBlockStmt.Eval(s)
	} else {
		return ite.elseBlockStmt.Eval(s)
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

func (ite IfThenElseStatement) Check(t *TypeState) bool {
	ty := ite.cond.Infer(t)
	if ty != TypeBool {
		return false
	}

	return ite.thenBlockStmt.Check(t) && ite.elseBlockStmt.Check(t)
}
