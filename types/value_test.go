package types

import (
	"imp/helper"
	"testing"
)

// TestIntValue tests the IntValue function

type TestIntValueCase struct {
	input     int
	want      Value
	compliant bool
}

var testIntValueTests = []TestIntValueCase{
	{0, Value{ValueType: ValueInt, IntValue: 0}, true},
	{0, Value{ValueType: ValueInt, IntValue: 1}, false},
}

func TestIntValue(t *testing.T) {
	for _, test := range testIntValueTests {
		if got := IntValue(test.input); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", helper.StructToJson(got), helper.StructToJson(test.want))
		}
	}
}

// TestBoolValue tests the BoolValue function

type TestBoolValueCase struct {
	input     bool
	want      Value
	compliant bool
}

var testBoolValueTests = []TestBoolValueCase{
	{true, Value{ValueType: ValueBool, BoolValue: true}, true},
	{true, Value{ValueType: ValueBool, BoolValue: false}, false},
}

func TestBoolValue(t *testing.T) {
	for _, test := range testBoolValueTests {
		if got := BoolValue(test.input); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", helper.StructToJson(got), helper.StructToJson(test.want))
		}
	}
}

// TestUndefinedValue tests the UndefinedValue function

type TestUndefinedValueCase struct {
	want      Value
	compliant bool
}

var testUndefinedValueTests = []TestUndefinedValueCase{
	{Value{ValueType: Undefined}, true},
	{Value{ValueType: ValueInt}, false},
}

func TestUndefinedValue(t *testing.T) {
	for _, test := range testUndefinedValueTests {
		if got := UndefinedValue(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", helper.StructToJson(got), helper.StructToJson(test.want))
		}
	}
}

// TestShowVal tests the ShowVal function

type TestShowValCase struct {
	input     Value
	want      string
	compliant bool
}

var testShowValTests = []TestShowValCase{
	{Value{ValueType: ValueInt, IntValue: 0}, "0", true},
	{Value{ValueType: ValueInt, IntValue: 1}, "1", true},
	{Value{ValueType: ValueBool, BoolValue: true}, "true", true},
	{Value{ValueType: ValueBool, BoolValue: false}, "false", true},
	{Value{ValueType: Undefined}, "Undefined", true},

	{Value{ValueType: ValueInt, IntValue: 0}, "1", false},
	{Value{ValueType: ValueInt, IntValue: 1}, "0", false},
	{Value{ValueType: ValueBool, BoolValue: true}, "false", false},
	{Value{ValueType: ValueBool, BoolValue: false}, "true", false},
	{Value{ValueType: Undefined}, "Defined", false},
}

func TestShowVal(t *testing.T) {
	for _, test := range testShowValTests {
		if got := ShowVal(test.input); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
