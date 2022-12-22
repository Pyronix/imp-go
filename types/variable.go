package types

type Variable string

func (x Variable) Pretty() string {
	return (string)(x)
}

func (x Variable) Eval(s ValueState) Value {
	val, ok := s[(string)(x)]
	if ok {
		return val
	} else {
		return Value{ValueType: Undefined} // variable does not exist yields undefined
	}
}

func (x Variable) Infer(t TypeState) Type {
	typ, ok := t[(string)(x)]
	if ok {
		return typ
	} else {
		return TypeIllTyped // variable does not exist yields illtyped
	}
}
