package statements

import . "imp/types"

type DeclarationStatement struct {
	lhs string
	rhs Expression
}

func Declaration(lhs string, rhs Expression) DeclarationStatement {
	return DeclarationStatement{lhs, rhs}
}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
func (decl DeclarationStatement) Eval(s ValueState) {
	s.Declare(decl.lhs, decl.rhs.Eval(s))
}

func (decl DeclarationStatement) Pretty() string {
	return decl.lhs + " := " + decl.rhs.Pretty()
}

func (decl DeclarationStatement) Check(t TypeState) bool {
	ty := decl.rhs.Infer(t)
	if ty == TypeIllTyped {
		return false
	}

	t.Declare(decl.lhs, ty)
	return true
}
