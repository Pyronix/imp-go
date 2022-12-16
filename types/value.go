package types

type Kind int

const (
	ValueInt  Kind = 0
	ValueBool Kind = 1
	Undefined Kind = 2
)

type Value struct {
	ValueType Kind
	IntValue  int
	BoolValue bool
}

// ValueState is a mapping from variable names to values
type ValueState map[string]Value

func IntValue(x int) Value {
	return Value{ValueType: ValueInt, IntValue: x}
}
func BoolValue(x bool) Value {
	return Value{ValueType: ValueBool, BoolValue: x}
}
func UndefinedValue() Value {
	return Value{ValueType: Undefined}
}

func ShowVal(v Value) string {
	var s string
	switch {
	case v.ValueType == ValueInt:
		s = NumberExpression(v.IntValue).Pretty()
	case v.ValueType == ValueBool:
		s = BoolExpression(v.BoolValue).Pretty()
	case v.ValueType == Undefined:
		s = "Undefined"
	}
	return s
}
