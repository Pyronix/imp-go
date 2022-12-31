package statements

import (
	"fmt"
	. "imp/types"
)

type WhileStatement struct {
	cond Expression
	stmt Statement
}

func While(cond Expression, stmt Statement) Statement {
	return WhileStatement{cond, stmt}
}

func (whi WhileStatement) Eval(s ValueState) {
	for {
		v := whi.cond.Eval(s)
		if v.ValueType != ValueBool {
			fmt.Printf("while Eval fail")
			break
		}
		if !v.BoolValue {
			break
		}
		whi.stmt.Eval(s)
	}
}

func (whi WhileStatement) Pretty() string {
	var x string
	x = "while "
	x += whi.cond.Pretty()
	x += " { "
	x += whi.stmt.Pretty()
	x += " }"

	return x
}

func (whi WhileStatement) Check(t TypeState) bool {
	ty := whi.cond.Infer(t)
	if ty != TypeBool {
		return false
	}

	return whi.stmt.Check(t)
}
