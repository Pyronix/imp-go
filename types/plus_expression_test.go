package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestPlus tests the Plus function

type TestPlusCase struct {
	input1    Expression
	input2    Expression
	want      PlusExpression
	compliant bool
}

var testPlusTests = []TestPlusCase{
	{NumberExpression(1), NumberExpression(1), PlusExpression{NumberExpression(1), NumberExpression(1)}, true},
	{NumberExpression(1), NumberExpression(1), PlusExpression{NumberExpression(10), NumberExpression(1)}, false},
}

func TestPlus(t *testing.T) {
	for _, test := range testPlusTests {
		if got := Plus(test.input1, test.input2); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

// TestPlusPretty tests the Pretty function

type TestPlusPrettyCase struct {
	input     PlusExpression
	want      string
	compliant bool
}

var testPlusPrettyTests = []TestPlusPrettyCase{
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, "1 + 1", true},

	{PlusExpression{BoolExpression(true), BoolExpression(true)}, "true + true", true},
	{PlusExpression{BoolExpression(false), BoolExpression(true)}, "false + true", true},

	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, "-1 + 1", false},
	{PlusExpression{BoolExpression(true), BoolExpression(true)}, "false + true", false},
}

func TestPlusPretty(t *testing.T) {
	for _, test := range testPlusPrettyTests {
		if got := test.input.Pretty(); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestPlusEval tests the Eval function

type TestPlusEvalCase struct {
	input     PlusExpression
	want      Value
	compliant bool
}

var testPlusEvalTests = []TestPlusEvalCase{
	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-2), true},

	{PlusExpression{NumberExpression(-1), BoolExpression(true)}, UndefinedValue(), true},
	{PlusExpression{BoolExpression(true), NumberExpression(-1)}, UndefinedValue(), true},

	{PlusExpression{BoolExpression(true), BoolExpression(true)}, UndefinedValue(), true},
	{PlusExpression{BoolExpression(true), BoolExpression(false)}, UndefinedValue(), true},

	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-1), false},
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, IntValue(2), false},

	{PlusExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(true), false},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(false), false},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, UndefinedValue(), false},

	{PlusExpression{NumberExpression(1), BoolExpression(false)}, IntValue(-1), false},
	{PlusExpression{NumberExpression(1), BoolExpression(false)}, BoolValue(false), false},
}

func TestPlusEval(t *testing.T) {
	for _, test := range testPlusEvalTests {
		if got := test.input.Eval(ValueState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

type TestPlusInferCase struct {
	input     PlusExpression
	want      Type
	compliant bool
}

var testPlusInferTests = []TestPlusInferCase{
	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, TypeInt, true},

	{PlusExpression{BoolExpression(true), BoolExpression(true)}, TypeIllTyped, true},
	{PlusExpression{BoolExpression(true), BoolExpression(false)}, TypeIllTyped, true},

	{PlusExpression{BoolExpression(false), NumberExpression(-1)}, TypeIllTyped, true},
	{PlusExpression{NumberExpression(-1), BoolExpression(true)}, TypeIllTyped, true},

	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, TypeBool, false},
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, TypeIllTyped, false},

	{PlusExpression{BoolExpression(false), NumberExpression(-1)}, TypeInt, false},
	{PlusExpression{NumberExpression(-1), BoolExpression(true)}, TypeBool, false},

	{PlusExpression{BoolExpression(false), BoolExpression(false)}, TypeInt, false},
	{PlusExpression{BoolExpression(true), BoolExpression(true)}, TypeBool, false},
}

// TestPlusInfer tests the Infer function

func TestPlusInfer(t *testing.T) {
	for _, test := range testPlusInferTests {
		if got := test.input.Infer(TypeState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
