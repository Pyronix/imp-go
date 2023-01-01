package types

type Type int

// TypeState is a mapping from variable names to types
type TypeState []map[string]Type

const (
	TypeIllTyped Type = 0
	TypeInt      Type = 1
	TypeBool     Type = 2
)

func ShowType(t Type) string {
	var s string
	switch {
	case t == TypeInt:
		s = "Int"
	case t == TypeBool:
		s = "BoolExpression"
	case t == TypeIllTyped:
		s = "Illtyped"
	}
	return s
}

func (ts TypeState) Declare(name string, newType Type) {
	ts[len(ts)-1][name] = newType
}

func (ts TypeState) Assign(name string, newType Type) {
	var variableExists bool = false

	for i := len(ts) - 1; i >= 0; i-- {
		if oldType, ok := ts[i][name]; ok {
			if oldType != newType {
				panic("Type mismatch")
			}
			variableExists = true
			ts[i][name] = newType
			break
		}
	}

	if !variableExists {
		panic("Variable " + name + " not declared")
	}
}

func PushTypeScope(ts *TypeState) {
	*ts = append(*ts, make(map[string]Type))
}

func PopTypeScope(ts *TypeState) {
	if len(*ts) > 1 {
		*ts = (*ts)[:len(*ts)-1]
	}
}

func (ts TypeState) GetCurrentTypeScope() map[string]Type {
	return ts[len(ts)-1]
}

func (ts TypeState) LookUpTypeByVariableName(name string) (Type, bool) {
	for i := len(ts) - 1; i >= 0; i-- {
		if typ, ok := ts[i][name]; ok {
			return typ, true
		}
	}
	return -1, false
}
