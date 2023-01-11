package statements

import (
	. "imp/types"
)

type BlockStatement struct {
	stmt Statement
}

func Block(x Statement) BlockStatement {
	return BlockStatement{x}
}

func (block BlockStatement) Pretty() string {
	return "{" + block.stmt.Pretty() + "}"
}

func (block BlockStatement) Eval(s *ValueState) {
	PushValueScope(s)
	block.stmt.Eval(s)
	PopValueScope(s)
}

func (block BlockStatement) Check(t *TypeState) bool {
	PushTypeScope(t)
	checkOk := block.stmt.Check(t)
	PopTypeScope(t)
	return checkOk
}
