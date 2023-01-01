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
type ValueState []map[string]Value

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

func (vs ValueState) Declare(name string, newValue Value) {
	vs[len(vs)-1][name] = newValue
}

func (vs ValueState) Assign(name string, newValue Value) {
	var variableExists bool = false

	for i := len(vs) - 1; i >= 0; i-- {
		if oldValue, ok := vs[i][name]; ok {
			if oldValue.ValueType != newValue.ValueType {
				panic("Type mismatch")
			}
			variableExists = true
			vs[i][name] = newValue
			break
		}
	}

	if !variableExists {
		panic("Variable " + name + " not declared")
	}
}

func PushValueScope(vs *ValueState) {
	*vs = append(*vs, make(map[string]Value))
}

func PopValueScope(vs *ValueState) {
	if len(*vs) > 1 {
		*vs = (*vs)[:len(*vs)-1]
	}
}

func (vs ValueState) GetCurrentValueScope() map[string]Value {
	return vs[len(vs)-1]
}

func (vs ValueState) LookUpValueByVariableName(name string) (Value, bool) {
	for i := len(vs) - 1; i >= 0; i-- {
		if value, ok := vs[i][name]; ok {
			return value, true
		}
	}
	return Value{}, false
}
