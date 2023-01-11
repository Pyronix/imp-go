package statements

import . "imp/types"

type ProgramStatement struct {
	stmt Statement
}

func Program(x Statement) ProgramStatement {
	return ProgramStatement{x}
}

func (prog ProgramStatement) Pretty() string {
	return prog.stmt.Pretty()
}

func (prog ProgramStatement) Eval(s *ValueState) Value {
	return prog.stmt.Eval(s)
}

func (prog ProgramStatement) Check(t *TypeState) bool {
	return prog.stmt.Check(t)
}
