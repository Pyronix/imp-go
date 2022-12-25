package statements

import . "imp/types"

type ProgramStatement BlockStatement

func (prog ProgramStatement) Pretty() string {
	return prog.Pretty()
}

func (prog ProgramStatement) Eval(s ValueState) {
	prog.Eval(s)
}

func (prog ProgramStatement) Check(t TypeState) bool {
	return prog.Check(t)
}
