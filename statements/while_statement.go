package statements

import (
	. "imp/types"
)

type WhileStatement struct {
	cond      Expression
	blockStmt BlockStatement
}

func While(cond Expression, blockStmt BlockStatement) Statement {
	return WhileStatement{cond, blockStmt}
}

func (whi WhileStatement) Eval(s *ValueState) {
	for {
		v := whi.cond.Eval(s)
		if v.ValueType != ValueBool {
			panic("while Eval fail")
		}
		if !v.BoolValue {
			break
		}
		whi.blockStmt.Eval(s)
	}
}

func (whi WhileStatement) Pretty() string {
	var x string
	x = "while "
	x += whi.cond.Pretty()
	x += " "
	x += whi.blockStmt.Pretty()

	return x
}

func (whi WhileStatement) Check(t *TypeState) bool {
	ty := whi.cond.Infer(t)
	if ty != TypeBool {
		return false
	}

	return whi.blockStmt.Check(t)
}
