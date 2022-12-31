package statements

import . "imp/types"

type ProgramStatement BlockStatement

func (prog ProgramStatement) Pretty() string {
	return prog.Statement.Pretty()
}

func (prog ProgramStatement) Eval(s ValueState) {
	prog.Statement.Eval(s)
}

func (prog ProgramStatement) Check(t TypeState) bool {
	return prog.Statement.Check(t)
}
