package types

import (
	. "imp/helper"
	"testing"
)

// TestOr tests the Equal function

type TestOrCase struct {
	input     Expression
	input2    Expression
	want      OrExpression
	compliant bool
}

var testOrTests = []TestOrCase{
	{BoolExpression(true), BoolExpression(true), OrExpression{BoolExpression(true), BoolExpression(true)}, true},
	{BoolExpression(false), BoolExpression(true), OrExpression{BoolExpression(false), BoolExpression(true)}, true},
	{NumberExpression(1), NumberExpression(1), OrExpression{NumberExpression(1), NumberExpression(1)}, true},
	{BoolExpression(true), NumberExpression(1), OrExpression{BoolExpression(true), NumberExpression(1)}, true},

	{BoolExpression(true), BoolExpression(true), OrExpression{BoolExpression(true), BoolExpression(false)}, false},
	{NumberExpression(1), NumberExpression(1), OrExpression{NumberExpression(1), NumberExpression(0)}, false},
	{BoolExpression(true), NumberExpression(1), OrExpression{BoolExpression(true), BoolExpression(true)}, false},
}

func TestOr(t *testing.T) {
	for _, test := range testOrTests {
		if got := Or(test.input, test.input2); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

// TestPretty tests the Pretty function

type TestOrPrettyCase struct {
	input     OrExpression
	want      string
	compliant bool
}

var testOrPrettyTests = []TestOrPrettyCase{
	{OrExpression{NumberExpression(-1), NumberExpression(-1)}, "-1||-1", true},

	{OrExpression{BoolExpression(true), BoolExpression(true)}, "true||true", true},
	{OrExpression{BoolExpression(false), BoolExpression(true)}, "false||true", true},
	{OrExpression{BoolExpression(true), BoolExpression(false)}, "true||false", true},
	{OrExpression{BoolExpression(false), BoolExpression(false)}, "false||false", true},

	{OrExpression{NumberExpression(-1), NumberExpression(-1)}, "-1||1", false},

	{OrExpression{BoolExpression(true), BoolExpression(true)}, "false||true", false},
	{OrExpression{BoolExpression(false), BoolExpression(true)}, "false||false", false},
}

func TestOrPretty(t *testing.T) {
	for _, test := range testOrPrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestEval tests the Eval function

type TestOrEvalCase struct {
	input     OrExpression
	want      Value
	compliant bool
}

var testOrEvalTests = []TestOrEvalCase{
	{OrExpression{BoolExpression(true), BoolExpression(true)}, BoolValue(true), true},
	{OrExpression{BoolExpression(false), BoolExpression(false)}, BoolValue(false), true},
	{OrExpression{BoolExpression(true), BoolExpression(false)}, BoolValue(true), true},
	{OrExpression{BoolExpression(false), BoolExpression(true)}, BoolValue(true), true},

	{OrExpression{NumberExpression(-1), BoolExpression(true)}, UndefinedValue(), true},
	{OrExpression{BoolExpression(true), NumberExpression(-1)}, BoolValue(true), true},

	{OrExpression{NumberExpression(-1), NumberExpression(-1)}, UndefinedValue(), true},

	{OrExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-1), false},
	{OrExpression{NumberExpression(0), NumberExpression(0)}, IntValue(2), false},

	{OrExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(true), false},
	{OrExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(false), false},
	{OrExpression{NumberExpression(1), NumberExpression(1)}, UndefinedValue(), false},

	{OrExpression{NumberExpression(1), BoolExpression(false)}, IntValue(-1), false},
	{OrExpression{BoolExpression(true), NumberExpression(1)}, BoolValue(true), true},
	{OrExpression{NumberExpression(1), BoolExpression(false)}, BoolValue(false), false},
	{OrExpression{BoolExpression(true), NumberExpression(1)}, BoolValue(true), true},
}

func TestOrEval(t *testing.T) {
	for _, test := range testOrEvalTests {
		if got := test.input.Eval(ValueState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

type TestOrInferCase struct {
	input     OrExpression
	want      Type
	compliant bool
}

var testOrInferTests = []TestOrInferCase{
	{OrExpression{BoolExpression(true), BoolExpression(true)}, TypeBool, true},
	{OrExpression{BoolExpression(false), BoolExpression(false)}, TypeBool, true},
	{OrExpression{BoolExpression(true), BoolExpression(false)}, TypeBool, true},
	{OrExpression{BoolExpression(false), BoolExpression(true)}, TypeBool, true},

	{OrExpression{NumberExpression(-1), BoolExpression(true)}, TypeIllTyped, true},
	{OrExpression{BoolExpression(true), NumberExpression(-1)}, TypeIllTyped, true},

	{OrExpression{NumberExpression(-1), NumberExpression(-1)}, TypeIllTyped, true},

	{OrExpression{NumberExpression(-1), NumberExpression(-1)}, TypeInt, false},
	{OrExpression{NumberExpression(0), NumberExpression(0)}, TypeInt, false},
}

// TestInfer tests the Infer function

func TestOrInfer(t *testing.T) {
	for _, test := range testOrInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
