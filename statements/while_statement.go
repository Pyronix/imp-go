package statements

import (
	"fmt"
	"imp/types"
)

type WhileStatement struct {
	cond types.Expression
	stmt Statement
}

func (whi WhileStatement) Eval(s types.ValueState) {
	// TODO: WIP
	v := whi.cond.Eval(s)

	if v.ValueType == types.ValueBool {

	} else {
		fmt.Printf("while Eval fail")
	}

	for v.BoolValue {
		v = whi.cond.Eval(s)

	}

}

func (whi WhileStatement) Pretty(s types.ValueState) string {
	var x string
	x = "while "
	x += whi.cond.Pretty()
	x += " { "
	x += whi.stmt.Pretty()
	x += " }"

	return x
}

func (whi WhileStatement) Check(t types.TypeState) bool {
	ty := whi.cond.Infer(t)
	if ty == types.TypeIllTyped {
		return false
	}

	return whi.stmt.Check(t)
}
