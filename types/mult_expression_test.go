package types

import (
	. "imp/helper"
	"testing"
)

// TestMult tests the Mult function

type TestMultCase struct {
	input1    Expression
	input2    Expression
	want      MultExpression
	compliant bool
}

var testMultTests = []TestMultCase{
	{NumberExpression(-1), NumberExpression(-1), MultExpression{NumberExpression(-1), NumberExpression(-1)}, true},
	{NumberExpression(0), NumberExpression(0), MultExpression{NumberExpression(0), NumberExpression(0)}, true},
	{NumberExpression(1), NumberExpression(1), MultExpression{NumberExpression(1), NumberExpression(1)}, true},

	{BoolExpression(true), BoolExpression(true), MultExpression{BoolExpression(true), BoolExpression(true)}, true},
	{BoolExpression(true), BoolExpression(false), MultExpression{BoolExpression(true), BoolExpression(false)}, true},
	{BoolExpression(false), BoolExpression(true), MultExpression{BoolExpression(false), BoolExpression(true)}, true},
	{BoolExpression(false), BoolExpression(false), MultExpression{BoolExpression(false), BoolExpression(false)}, true},

	{NumberExpression(-1), NumberExpression(-1), MultExpression{NumberExpression(1), NumberExpression(-1)}, false},
	{NumberExpression(0), NumberExpression(0), MultExpression{NumberExpression(-1), NumberExpression(0)}, false},
	{NumberExpression(1), NumberExpression(1), MultExpression{NumberExpression(1), NumberExpression(0)}, false},

	{BoolExpression(true), BoolExpression(true), MultExpression{BoolExpression(false), BoolExpression(true)}, false},
	{BoolExpression(true), BoolExpression(false), MultExpression{BoolExpression(false), BoolExpression(false)}, false},
	{BoolExpression(false), BoolExpression(true), MultExpression{BoolExpression(true), BoolExpression(true)}, false},
	{BoolExpression(false), BoolExpression(false), MultExpression{BoolExpression(true), BoolExpression(false)}, false},
}

func TestMultNumber(t *testing.T) {
	for _, test := range testMultTests {
		if got := Mult(test.input1, test.input2); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestPretty tests the Pretty function

type TestMultPrettyCase struct {
	input     MultExpression
	want      string
	compliant bool
}

var testMultPrettyTests = []TestMultPrettyCase{
	{MultExpression{NumberExpression(-1), NumberExpression(-1)}, "-1*-1", true},
	{MultExpression{NumberExpression(0), NumberExpression(0)}, "0*0", true},
	{MultExpression{NumberExpression(1), NumberExpression(1)}, "1*1", true},

	{MultExpression{BoolExpression(true), BoolExpression(true)}, "true*true", true},
	{MultExpression{BoolExpression(false), BoolExpression(true)}, "false*true", true},
	{MultExpression{BoolExpression(true), BoolExpression(false)}, "true*false", true},
	{MultExpression{BoolExpression(false), BoolExpression(false)}, "false*false", true},

	{MultExpression{NumberExpression(-1), NumberExpression(-1)}, "-1*1", false},
	{MultExpression{NumberExpression(0), NumberExpression(0)}, "-1*0", false},
	{MultExpression{NumberExpression(1), NumberExpression(1)}, "0*0", false},

	{MultExpression{BoolExpression(true), BoolExpression(true)}, "false*true", false},
	{MultExpression{BoolExpression(false), BoolExpression(true)}, "false*false", false},
	{MultExpression{BoolExpression(true), BoolExpression(false)}, "true*true", false},
	{MultExpression{BoolExpression(false), BoolExpression(false)}, "true*true", false},
}

func TestMultPretty(t *testing.T) {
	for _, test := range testMultPrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestEval tests the Eval function

type TestMultEvalCase struct {
	input     MultExpression
	want      Value
	compliant bool
}

var testMultEvalTests = []TestMultEvalCase{
	{MultExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(1), true},
	{MultExpression{NumberExpression(0), NumberExpression(0)}, IntValue(0), true},
	{MultExpression{NumberExpression(2), NumberExpression(2)}, IntValue(4), true},

	{MultExpression{NumberExpression(-1), BoolExpression(true)}, UndefinedValue(), true},
	{MultExpression{BoolExpression(true), NumberExpression(-1)}, UndefinedValue(), true},

	{MultExpression{BoolExpression(true), BoolExpression(true)}, UndefinedValue(), true},
	{MultExpression{BoolExpression(true), BoolExpression(false)}, UndefinedValue(), true},
	{MultExpression{BoolExpression(false), BoolExpression(true)}, UndefinedValue(), true},
	{MultExpression{BoolExpression(false), BoolExpression(false)}, UndefinedValue(), true},

	{MultExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-1), false},
	{MultExpression{NumberExpression(0), NumberExpression(0)}, IntValue(2), false},

	{MultExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(true), false},
	{MultExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(false), false},
	{MultExpression{NumberExpression(1), NumberExpression(1)}, UndefinedValue(), false},

	{MultExpression{NumberExpression(1), BoolExpression(false)}, IntValue(-1), false},
	{MultExpression{BoolExpression(true), NumberExpression(1)}, IntValue(-1), false},
	{MultExpression{NumberExpression(1), BoolExpression(false)}, BoolValue(false), false},
	{MultExpression{BoolExpression(true), NumberExpression(1)}, BoolValue(true), false},
}

func TestMultPrettyEval(t *testing.T) {
	for _, test := range testMultEvalTests {
		if got := test.input.Eval(ValueState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

type TestMultInferCase struct {
	input     MultExpression
	want      Type
	compliant bool
}

var testMultInferTests = []TestMultInferCase{
	{MultExpression{NumberExpression(-1), NumberExpression(-1)}, TypeInt, true},
	{MultExpression{NumberExpression(0), NumberExpression(0)}, TypeInt, true},
	{MultExpression{NumberExpression(1), NumberExpression(1)}, TypeInt, true},

	{MultExpression{BoolExpression(true), BoolExpression(true)}, TypeIllTyped, true},
	{MultExpression{BoolExpression(true), BoolExpression(false)}, TypeIllTyped, true},
	{MultExpression{BoolExpression(false), BoolExpression(true)}, TypeIllTyped, true},
	{MultExpression{BoolExpression(false), BoolExpression(false)}, TypeIllTyped, true},

	{MultExpression{BoolExpression(false), NumberExpression(-1)}, TypeIllTyped, true},
	{MultExpression{NumberExpression(-1), BoolExpression(true)}, TypeIllTyped, true},

	{MultExpression{NumberExpression(-1), NumberExpression(-1)}, TypeBool, false},
	{MultExpression{NumberExpression(0), NumberExpression(0)}, TypeIllTyped, false},

	{MultExpression{BoolExpression(false), NumberExpression(-1)}, TypeInt, false},
	{MultExpression{NumberExpression(-1), BoolExpression(true)}, TypeBool, false},

	{MultExpression{BoolExpression(false), BoolExpression(false)}, TypeInt, false},
	{MultExpression{BoolExpression(true), BoolExpression(true)}, TypeBool, false},
}

// TestInfer tests the Infer function

func TestMultInfer(t *testing.T) {
	for _, test := range testMultInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
