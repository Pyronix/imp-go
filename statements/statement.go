package statements

import "imp/types"

type Statement interface {
	Pretty() string
	Eval(s types.ValueState)
	Check(t types.TypeState) bool
}
