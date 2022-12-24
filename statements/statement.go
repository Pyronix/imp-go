package statements

import . "imp/types"

type Statement interface {
	Pretty() string
	Eval(s ValueState)
	Check(t TypeState) bool
}
