package types

import (
	. "imp/helper"
	"testing"
)

// TestPretty tests the Pretty function

type TestPlusPrettyCase struct {
	input     PlusExpression
	want      string
	compliant bool
}

var testPlusPrettyTests = []TestPlusPrettyCase{
	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, "-1+-1", true},
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, "0+0", true},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, "1+1", true},

	{PlusExpression{BoolExpression(true), BoolExpression(true)}, "true+true", true},
	{PlusExpression{BoolExpression(false), BoolExpression(true)}, "false+true", true},
	{PlusExpression{BoolExpression(true), BoolExpression(false)}, "true+false", true},
	{PlusExpression{BoolExpression(false), BoolExpression(false)}, "false+false", true},

	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, "-1+1", false},
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, "-1+0", false},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, "0+0", false},

	{PlusExpression{BoolExpression(true), BoolExpression(true)}, "false+true", false},
	{PlusExpression{BoolExpression(false), BoolExpression(true)}, "false+false", false},
	{PlusExpression{BoolExpression(true), BoolExpression(false)}, "true+true", false},
	{PlusExpression{BoolExpression(false), BoolExpression(false)}, "true+true", false},
}

func TestPlusPretty(t *testing.T) {
	for _, test := range testPlusPrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestEval tests the Eval function

type TestPlusEvalCase struct {
	input     PlusExpression
	want      Value
	compliant bool
}

var testPlusEvalTests = []TestPlusEvalCase{
	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-2), true},
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, IntValue(0), true},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, IntValue(2), true},

	{PlusExpression{NumberExpression(-1), BoolExpression(true)}, UndefinedValue(), true},
	{PlusExpression{BoolExpression(true), NumberExpression(-1)}, UndefinedValue(), true},

	{PlusExpression{BoolExpression(true), BoolExpression(true)}, UndefinedValue(), true},
	{PlusExpression{BoolExpression(true), BoolExpression(false)}, UndefinedValue(), true},
	{PlusExpression{BoolExpression(false), BoolExpression(true)}, UndefinedValue(), true},
	{PlusExpression{BoolExpression(false), BoolExpression(false)}, UndefinedValue(), true},

	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-1), false},
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, IntValue(2), false},

	{PlusExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(true), false},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(false), false},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, UndefinedValue(), false},

	{PlusExpression{NumberExpression(1), BoolExpression(false)}, IntValue(-1), false},
	{PlusExpression{BoolExpression(true), NumberExpression(1)}, IntValue(-1), false},
	{PlusExpression{NumberExpression(1), BoolExpression(false)}, BoolValue(false), false},
	{PlusExpression{BoolExpression(true), NumberExpression(1)}, BoolValue(true), false},
}

func TestPlusEval(t *testing.T) {
	for _, test := range testPlusEvalTests {
		if got := test.input.Eval(ValueState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
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
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, TypeInt, true},
	{PlusExpression{NumberExpression(1), NumberExpression(1)}, TypeInt, true},

	{PlusExpression{BoolExpression(true), BoolExpression(true)}, TypeIllTyped, true},
	{PlusExpression{BoolExpression(true), BoolExpression(false)}, TypeIllTyped, true},
	{PlusExpression{BoolExpression(false), BoolExpression(true)}, TypeIllTyped, true},
	{PlusExpression{BoolExpression(false), BoolExpression(false)}, TypeIllTyped, true},

	{PlusExpression{BoolExpression(false), NumberExpression(-1)}, TypeIllTyped, true},
	{PlusExpression{NumberExpression(-1), BoolExpression(true)}, TypeIllTyped, true},

	{PlusExpression{NumberExpression(-1), NumberExpression(-1)}, TypeBool, false},
	{PlusExpression{NumberExpression(0), NumberExpression(0)}, TypeIllTyped, false},

	{PlusExpression{BoolExpression(false), NumberExpression(-1)}, TypeInt, false},
	{PlusExpression{NumberExpression(-1), BoolExpression(true)}, TypeBool, false},

	{PlusExpression{BoolExpression(false), BoolExpression(false)}, TypeInt, false},
	{PlusExpression{BoolExpression(true), BoolExpression(true)}, TypeBool, false},
}

// TestInfer tests the Infer function

func TestPlusInfer(t *testing.T) {
	for _, test := range testPlusInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
