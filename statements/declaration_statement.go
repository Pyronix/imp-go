package statements

import "imp/types"

type DeclarationStatement struct {
	lhs string
	rhs types.Expression
}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
func (decl DeclarationStatement) Eval(s types.ValueState) {
	v := decl.rhs.Eval(s)
	x := (string)(decl.lhs)
	s[x] = v
}

func (decl DeclarationStatement) Pretty() string {
	return decl.lhs + " := " + decl.rhs.Pretty()
}

func (decl DeclarationStatement) Check(t types.TypeState) bool {
	ty := decl.rhs.Infer(t)
	if ty == types.TypeIllTyped {
		return false
	}

	x := (string)(decl.lhs)
	t[x] = ty
	return true
}
