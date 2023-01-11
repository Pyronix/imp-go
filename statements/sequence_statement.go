package statements

import . "imp/types"

type SequenceStatement [2]Statement

func Sequence(x, y Statement) Statement {
	return SequenceStatement{x, y}
}

func (stmt SequenceStatement) Pretty() string {
	return stmt[0].Pretty() + "; " + stmt[1].Pretty()
}

func (stmt SequenceStatement) Eval(s *ValueState) {
	stmt[0].Eval(s)
	stmt[1].Eval(s)
}

func (stmt SequenceStatement) Check(t *TypeState) bool {
	return stmt[0].Check(t) && stmt[1].Check(t)
}
