package statements

import "imp/types"

type Sequence [2]Statement

func (stmt Sequence) Pretty() string {
	return stmt[0].Pretty() + "; " + stmt[1].Pretty()
}

func (stmt Sequence) Eval(s types.ValueState) {
	stmt[0].Eval(s)
	stmt[1].Eval(s)
}

func (stmt Sequence) Check(t types.TypeState) bool {
	if !stmt[0].Check(t) {
		return false
	}
	return stmt[1].Check(t)
}
