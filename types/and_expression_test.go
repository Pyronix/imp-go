package types

import (
	. "imp/helper"
	"testing"
)

// TestPretty tests the Pretty function

type TestAndPrettyCase struct {
	input     AndExpression
	want      string
	compliant bool
}

var testAndPrettyTests = []TestAndPrettyCase{
	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, "-1&&-1", true},

	{AndExpression{BoolExpression(true), BoolExpression(true)}, "true&&true", true},
	{AndExpression{BoolExpression(false), BoolExpression(true)}, "false&&true", true},
	{AndExpression{BoolExpression(true), BoolExpression(false)}, "true&&false", true},
	{AndExpression{BoolExpression(false), BoolExpression(false)}, "false&&false", true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, "-1&&1", false},

	{AndExpression{BoolExpression(true), BoolExpression(true)}, "false&&true", false},
	{AndExpression{BoolExpression(false), BoolExpression(true)}, "false&&false", false},
}

func TestAndPretty(t *testing.T) {
	for _, test := range testAndPrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestEval tests the Eval function

type TestAndEvalCase struct {
	input     AndExpression
	want      Value
	compliant bool
}

var testAndEvalTests = []TestAndEvalCase{
	{AndExpression{BoolExpression(true), BoolExpression(true)}, BoolValue(true), true},
	{AndExpression{BoolExpression(false), BoolExpression(false)}, BoolValue(false), true},
	{AndExpression{BoolExpression(true), BoolExpression(false)}, BoolValue(false), true},
	{AndExpression{BoolExpression(false), BoolExpression(true)}, BoolValue(false), true},

	{AndExpression{NumberExpression(-1), BoolExpression(true)}, UndefinedValue(), true},
	{AndExpression{BoolExpression(true), NumberExpression(-1)}, UndefinedValue(), true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, UndefinedValue(), true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-1), false},
	{AndExpression{NumberExpression(0), NumberExpression(0)}, IntValue(2), false},

	{AndExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(true), false},
	{AndExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(false), false},
	{AndExpression{NumberExpression(1), NumberExpression(1)}, UndefinedValue(), false},

	{AndExpression{NumberExpression(1), BoolExpression(false)}, IntValue(-1), false},
	{AndExpression{BoolExpression(true), NumberExpression(1)}, IntValue(-1), false},
	{AndExpression{NumberExpression(1), BoolExpression(false)}, BoolValue(false), false},
	{AndExpression{BoolExpression(true), NumberExpression(1)}, BoolValue(true), false},
}

func TestAndEval(t *testing.T) {
	for _, test := range testAndEvalTests {
		if got := test.input.Eval(ValueState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

type TestAndInferCase struct {
	input     AndExpression
	want      Type
	compliant bool
}

var testAndInferTests = []TestAndInferCase{
	{AndExpression{BoolExpression(true), BoolExpression(true)}, TypeBool, true},
	{AndExpression{BoolExpression(false), BoolExpression(false)}, TypeBool, true},
	{AndExpression{BoolExpression(true), BoolExpression(false)}, TypeBool, true},
	{AndExpression{BoolExpression(false), BoolExpression(true)}, TypeBool, true},

	{AndExpression{NumberExpression(-1), BoolExpression(true)}, TypeIllTyped, true},
	{AndExpression{BoolExpression(true), NumberExpression(-1)}, TypeIllTyped, true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, TypeIllTyped, true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, TypeInt, false},
	{AndExpression{NumberExpression(0), NumberExpression(0)}, TypeInt, false},
}

// TestInfer tests the Infer function

func TestAndInfer(t *testing.T) {
	for _, test := range testAndInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}