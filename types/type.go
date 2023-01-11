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

func (ts *TypeState) Declare(name string, newType Type) {
	(*ts)[len(*ts)-1][name] = newType
}

func (ts *TypeState) Assign(name string, newType Type) {
	for i := len(*ts) - 1; i >= 0; i-- {
		if oldType, ok := (*ts)[i][name]; ok {
			if oldType != newType {
				panic("Type mismatch")
			}

			(*ts)[i][name] = newType

			return
		}
	}

	panic("Variable " + name + " not declared")
}

func PushTypeScope(ts *TypeState) {
	if len(*ts) == 0 {
		*ts = append(*ts, make(map[string]Type))
	}
	*ts = append(*ts, make(map[string]Type))
}

func PopTypeScope(ts *TypeState) {
	if len(*ts) < 2 {
		panic("Cannot unscope global scope")
	}
	*ts = (*ts)[:len(*ts)-1]
}

func (ts *TypeState) GetCurrentTypeScope() map[string]Type {
	if len(*ts) == 0 {
		*ts = append(*ts, make(map[string]Type))
	}
	return (*ts)[len(*ts)-1]
}

func (ts *TypeState) LookUpTypeByVariableName(name string) (Type, bool) {
	for i := len(*ts) - 1; i >= 0; i-- {
		if typ, ok := (*ts)[i][name]; ok {
			return typ, true
		}
	}
	return TypeIllTyped, false
}
