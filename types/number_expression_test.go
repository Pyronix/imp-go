package types

import (
	. "imp/helper"
	"testing"
)

// TestNumber tests the Number function

type TestNumberCase struct {
	input     int
	want      NumberExpression
	compliant bool
}

var testNumberTests = []TestNumberCase{
	{-1, NumberExpression(-1), true},
	{0, NumberExpression(0), true},
	{1, NumberExpression(1), true},
	{-1, NumberExpression(-10), false},
	{0, NumberExpression(999), false},
	{1, NumberExpression(10), false},
}

func TestNumber(t *testing.T) {
	for _, test := range testNumberTests {
		if got := Number(test.input); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestPretty tests the Pretty function

type TestPrettyCase struct {
	input     NumberExpression
	want      string
	compliant bool
}

var testPrettyTests = []TestPrettyCase{
	{NumberExpression(-1), "-1", true},
	{NumberExpression(0), "0", true},
	{NumberExpression(1), "1", true},
	{NumberExpression(-1), "-10", false},
	{NumberExpression(0), "999", false},
	{NumberExpression(1), "10", false},
}

func TestPretty(t *testing.T) {
	for _, test := range testPrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestEval tests the Eval function

type TestEvalCase struct {
	input     NumberExpression
	want      Value
	compliant bool
}

var testEvalTests = []TestEvalCase{
	{NumberExpression(-1), IntValue(-1), true},
	{NumberExpression(0), IntValue(0), true},
	{NumberExpression(1), IntValue(1), true},
	{NumberExpression(-1), IntValue(-10), false},
	{NumberExpression(0), IntValue(999), false},
	{NumberExpression(1), IntValue(10), false},
}

func TestEval(t *testing.T) {
	for _, test := range testEvalTests {
		if got := test.input.Eval(ValueState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

type TestInferCase struct {
	input     NumberExpression
	want      Type
	compliant bool
}

var testInferTests = []TestInferCase{
	{NumberExpression(-1), TypeInt, true},
	{NumberExpression(0), TypeInt, true},
	{NumberExpression(1), TypeInt, true},
	{NumberExpression(-1), TypeBool, false},
	{NumberExpression(0), TypeBool, false},
	{NumberExpression(1), TypeBool, false},
}

// TestInfer tests the Infer function

func TestInfer(t *testing.T) {
	for _, test := range testInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
