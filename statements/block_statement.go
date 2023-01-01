package statements

import . "imp/types"

type BlockStatement struct {
	Statement
}

func Block(x Statement) BlockStatement {
	return BlockStatement{x}
}

func (block BlockStatement) Pretty() string {
	return "{" + block.Statement.Pretty() + "}"
}

func (block BlockStatement) Eval(s ValueState) {
	block.Statement.Eval(s)
}

func (block BlockStatement) Check(t TypeState) bool {
	return block.Statement.Check(t)
}
