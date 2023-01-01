package statements

import . "imp/types"

type ProgramStatement struct {
	blockStmt BlockStatement
}

func Program(x BlockStatement) ProgramStatement {
	return ProgramStatement{x}
}

func (prog ProgramStatement) Pretty() string {
	return prog.blockStmt.Pretty()
}

func (prog ProgramStatement) Eval(s ValueState) {
	prog.blockStmt.Eval(s)
}

func (prog ProgramStatement) Check(t TypeState) bool {
	return prog.blockStmt.Check(t)
}
