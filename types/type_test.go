package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestShowType tests the ShowType function

type TestShowTypeCase struct {
	input     Type
	want      string
	compliant bool
}

var testShowTypeTests = []TestShowTypeCase{
	{TypeIllTyped, "Illtyped", true},
	{TypeInt, "Int", true},
	{TypeBool, "BoolExpression", true},

	{TypeIllTyped, "Int", false},
	{TypeInt, "BoolExpression", false},
	{TypeBool, "Illtyped", false},
}

func TestShowType(t *testing.T) {
	for _, test := range testShowTypeTests {
		if got := ShowType(test.input); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
