package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestEqual tests the Equal function

type TestEqualCase struct {
	input     Expression
	input2    Expression
	want      EqualityExpression
	compliant bool
}

var testEqualTests = []TestEqualCase{
	{BoolExpression(true), BoolExpression(true), EqualityExpression{BoolExpression(true), BoolExpression(true)}, true},
	{BoolExpression(false), BoolExpression(true), EqualityExpression{BoolExpression(false), BoolExpression(true)}, true},
	{NumberExpression(1), NumberExpression(1), EqualityExpression{NumberExpression(1), NumberExpression(1)}, true},

	{BoolExpression(true), NumberExpression(1), EqualityExpression{BoolExpression(true), NumberExpression(1)}, true},

	{BoolExpression(true), BoolExpression(true), EqualityExpression{BoolExpression(true), BoolExpression(false)}, false},
	{NumberExpression(1), NumberExpression(1), EqualityExpression{NumberExpression(1), NumberExpression(0)}, false},

	{BoolExpression(true), NumberExpression(1), EqualityExpression{BoolExpression(true), BoolExpression(true)}, false},
	{BoolExpression(true), NumberExpression(-1), EqualityExpression{BoolExpression(true), BoolExpression(false)}, false},
}

func TestEqual(t *testing.T) {
	for _, test := range testEqualTests {
		if got := Equal(test.input, test.input2); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestEqualityPretty tests the Pretty function

type TestEqualityPrettyCase struct {
	input     EqualityExpression
	want      string
	compliant bool
}

var testEqualityPrettyTests = []TestEqualityPrettyCase{
	{EqualityExpression{NumberExpression(-1), NumberExpression(-1)}, "-1 == -1", true},
	{EqualityExpression{NumberExpression(-1), NumberExpression(1)}, "-1 == 1", true},
	{EqualityExpression{NumberExpression(1), NumberExpression(-1)}, "1 == -1", true},
	{EqualityExpression{BoolExpression(true), NumberExpression(0)}, "true == 0", true},
	{EqualityExpression{NumberExpression(0), BoolExpression(true)}, "0 == true", true},
	{EqualityExpression{BoolExpression(true), BoolExpression(false)}, "true == false", true},
	{EqualityExpression{BoolExpression(true), BoolExpression(true)}, "???", false},
	{EqualityExpression{NumberExpression(-1), NumberExpression(-1)}, "???", false},
	{EqualityExpression{NumberExpression(0), BoolExpression(true)}, "???", false},
}

func TestEqualityPretty(t *testing.T) {
	for _, test := range testEqualityPrettyTests {
		if got := test.input.Pretty(); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestEqualityEval tests the Eval function

type TestEqualityEvalCase struct {
	input     Expression
	want      Value
	compliant bool
}

var TestEqualityEvalTests = []TestEqualityEvalCase{
	{EqualityExpression{BoolExpression(true), BoolExpression(true)}, BoolValue(true), true},
	{EqualityExpression{BoolExpression(true), BoolExpression(false)}, BoolValue(false), true},

	{EqualityExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(true), true},
	{EqualityExpression{NumberExpression(1), NumberExpression(-1)}, BoolValue(false), true},
	{EqualityExpression{BoolExpression(true), NumberExpression(1)}, UndefinedValue(), true},
	{EqualityExpression{NumberExpression(1), BoolExpression(true)}, UndefinedValue(), true},

	{EqualityExpression{BoolExpression(true), NumberExpression(1)}, BoolValue(true), false},
	{EqualityExpression{BoolExpression(false), NumberExpression(1)}, IntValue(1), false},
	{EqualityExpression{BoolExpression(false), NumberExpression(1)}, BoolValue(false), false},

	{EqualityExpression{BoolExpression(true), BoolExpression(true)}, UndefinedValue(), false},
	{EqualityExpression{NumberExpression(1), NumberExpression(1)}, UndefinedValue(), false},
}

func TestEqualityEval(t *testing.T) {
	for _, test := range TestEqualityEvalTests {
		if got := test.input.Eval(&ValueState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

type TestEqualityInferCase struct {
	input     EqualityExpression
	want      Type
	compliant bool
}

var TestEqualityInferTests = []TestEqualityInferCase{
	{EqualityExpression{NumberExpression(-1), NumberExpression(-1)}, TypeBool, true},
	{EqualityExpression{BoolExpression(true), BoolExpression(true)}, TypeBool, true},
	{EqualityExpression{NumberExpression(-1), BoolExpression(true)}, TypeIllTyped, true},

	{EqualityExpression{NumberExpression(-1), BoolExpression(true)}, TypeInt, false},
	{EqualityExpression{NumberExpression(-1), BoolExpression(true)}, TypeBool, false},
	{EqualityExpression{BoolExpression(true), BoolExpression(true)}, TypeIllTyped, false},
	{EqualityExpression{BoolExpression(true), BoolExpression(true)}, TypeInt, false},
}

// TestEqualityInfer tests the Infer function

func TestEqualityInfer(t *testing.T) {
	for _, test := range TestEqualityInferTests {
		if got := test.input.Infer(&TypeState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
