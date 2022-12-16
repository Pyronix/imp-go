package types

type Variable string

func (x Variable) Pretty() string {
	return (string)(x)
}

func (x Variable) Infer(t TypeState) Type {
	y := (string)(x)
	ty, ok := t[y]
	if ok {
		return ty
	} else {
		return TypeIllTyped // variable does not exist yields illtyped
	}
}
