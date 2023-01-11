package types

type Expression interface {
	Pretty() string
	Eval(s *ValueState) Value
	Infer(t *TypeState) Type
}
