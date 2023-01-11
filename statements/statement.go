package statements

import (
	. "imp/types"
)

type Statement interface {
	Pretty() string
	Eval(s *ValueState) Value
	Check(t *TypeState) bool
}
