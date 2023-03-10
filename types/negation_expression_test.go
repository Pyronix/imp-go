package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestNegation tests the Neegation function

type TestNegationCase struct {
	input     Expression
	want      Expression
	compliant bool
}

var testNegationTests = []TestNegationCase{
	{NumberExpression(1), NegationExpression{NumberExpression(1)}, true},
	{NumberExpression(1), NegationExpression{NumberExpression(-1)}, false},
}

func TestNegation(t *testing.T) {
	for _, test := range testNegationTests {
		if got := Negation(test.input); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestNegationPretty tests the Pretty function

type TestNegationPrettyCase struct {
	input     NegationExpression
	want      string
	compliant bool
}

var testNegationPrettyTests = []TestNegationPrettyCase{
	{NegationExpression{NumberExpression(1)}, "!1", true},
	{NegationExpression{NumberExpression(1)}, "!-1", false},
}

func TestNegationPretty(t *testing.T) {
	for _, test := range testNegationPrettyTests {
		if got := test.input.Pretty(); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestNegationEval tests the Eval function

type TestNegationEvalCase struct {
	input     NegationExpression
	want      Value
	compliant bool
}

var testNegationEvalTests = []TestNegationEvalCase{
	{NegationExpression{BoolExpression(true)}, BoolValue(false), true},
	{NegationExpression{BoolExpression(false)}, BoolValue(true), true},
	{NegationExpression{NumberExpression(-1)}, UndefinedValue(), true},
	{NegationExpression{BoolExpression(true)}, BoolValue(true), false},
	{NegationExpression{BoolExpression(false)}, BoolValue(false), false},
	{NegationExpression{NumberExpression(-1)}, IntValue(-1), false},
}

func TestNegationEval(t *testing.T) {
	for _, test := range testNegationEvalTests {
		if got := test.input.Eval(&ValueState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestNegationInfer tests the Infer function

type TestNegationInferCase struct {
	input     NegationExpression
	want      Type
	compliant bool
}

var testNegationInferTests = []TestNegationInferCase{
	{NegationExpression{BoolExpression(true)}, TypeBool, true},
	{NegationExpression{NumberExpression(1)}, TypeIllTyped, true},
	{NegationExpression{NumberExpression(1)}, TypeInt, false},
}

func TestNegationInfer(t *testing.T) {
	for _, test := range testNegationInferTests {
		if got := test.input.Infer(&TypeState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
