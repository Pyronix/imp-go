package types

import (
	. "imp/helper"
	"testing"
)

// TestEqual tests the Equal function

type TestEqualCase struct {
	input     NumberExpression
	input2    NumberExpression
	want      EqualityExpression
	compliant bool
}

var testEqualTests = []TestEqualCase{
	{NumberExpression(-1), NumberExpression(-1), EqualityExpression{NumberExpression(-1), NumberExpression(-1)}, true},
	{NumberExpression(0), NumberExpression(0), EqualityExpression{NumberExpression(0), NumberExpression(0)}, true},
	{NumberExpression(1), NumberExpression(1), EqualityExpression{NumberExpression(1), NumberExpression(1)}, true},
	{NumberExpression(-1), NumberExpression(-10), EqualityExpression{NumberExpression(-1), NumberExpression(-10)}, false},
	{NumberExpression(0), NumberExpression(999), EqualityExpression{NumberExpression(0), NumberExpression(999)}, false},
	{NumberExpression(1), NumberExpression(10), EqualityExpression{NumberExpression(1), NumberExpression(10)}, false},
}

func TestEqual(t *testing.T) {
	for _, test := range testEqualTests {
		if got := Equal(test.input, test.input2); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestPretty tests the Pretty function

type TestEqualityPrettyCase struct {
	input     EqualityExpression
	want      string
	compliant bool
}

var testEqualityPrettyTests = []TestEqualityPrettyCase{
	{EqualityExpression{NumberExpression(-1), NumberExpression(-1)}, "(-1 == -1)", true},
	{EqualityExpression{NumberExpression(-1), NumberExpression(1)}, "(-1 == 1)", true},
	{EqualityExpression{NumberExpression(1), NumberExpression(-1)}, "(1 == -1)", true},
	{EqualityExpression{BoolExpression(true), NumberExpression(0)}, "(true == 0)", true},
	{EqualityExpression{NumberExpression(0), BoolExpression(true)}, "(0 == true)", true},
	{EqualityExpression{BoolExpression(true), BoolExpression(false)}, "(true == false)", true},
	{EqualityExpression{BoolExpression(true), BoolExpression(true)}, "???", false},
	{EqualityExpression{NumberExpression(-1), NumberExpression(-1)}, "???", false},
	{EqualityExpression{NumberExpression(0), BoolExpression(true)}, "???", false},
}

func TestEqualityPretty(t *testing.T) {
	for _, test := range testEqualityPrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestEqualityEval tests the Eval function

type TestEqualityEvalCase struct {
	input     NumberExpression
	want      Value
	compliant bool
}

var TestEqualityEvalTests = []TestEqualityEvalCase{
	{NumberExpression(-1), IntValue(-1), true},
	{NumberExpression(0), IntValue(0), true},
	{NumberExpression(1), IntValue(1), true},
	{NumberExpression(-1), IntValue(-10), false},
	{NumberExpression(0), IntValue(999), false},
	{NumberExpression(1), IntValue(10), false},
}

func TestEqualityEval(t *testing.T) {
	for _, test := range TestEqualityEvalTests {
		if got := test.input.Eval(ValueState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

type TestEqualityInferCase struct {
	input     NumberExpression
	want      Type
	compliant bool
}

var TestEqualityInferTests = []TestEqualityInferCase{
	{NumberExpression(-1), TypeInt, true},
	{NumberExpression(0), TypeInt, true},
	{NumberExpression(1), TypeInt, true},
	{NumberExpression(-1), TypeBool, false},
	{NumberExpression(0), TypeBool, false},
	{NumberExpression(1), TypeBool, false},
}

// TestEqualityInfer tests the Infer function

func TestEqualityInfer(t *testing.T) {
	for _, test := range TestEqualityInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
