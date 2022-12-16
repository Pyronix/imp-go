package types

type Type int

// TypeState is a mapping from variable names to types
type TypeState map[string]Type

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
