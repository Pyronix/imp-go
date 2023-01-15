package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestAnd tests the And function

type TestAndCase struct {
	input1    Expression
	input2    Expression
	want      AndExpression
	compliant bool
}

var testAndTests = []TestAndCase{
	{BoolExpression(true), BoolExpression(true), AndExpression{BoolExpression(true), BoolExpression(true)}, true},
	{BoolExpression(false), BoolExpression(true), AndExpression{BoolExpression(false), BoolExpression(true)}, true},

	{BoolExpression(true), NumberExpression(1), AndExpression{BoolExpression(true), NumberExpression(1)}, true},
	{NumberExpression(-1), NumberExpression(1), AndExpression{NumberExpression(-1), NumberExpression(1)}, true},

	{BoolExpression(true), BoolExpression(true), AndExpression{BoolExpression(false), BoolExpression(true)}, false},

	{BoolExpression(true), NumberExpression(1), AndExpression{BoolExpression(true), BoolExpression(true)}, false},
	{NumberExpression(-1), BoolExpression(true), AndExpression{BoolExpression(false), BoolExpression(true)}, false},
}

func TestAnd(t *testing.T) {
	for _, test := range testAndTests {
		if got := And(test.input1, test.input2); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestAndPretty tests the Pretty function

type TestAndPrettyCase struct {
	input     AndExpression
	want      string
	compliant bool
}

var testAndPrettyTests = []TestAndPrettyCase{
	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, "-1 && -1", true},

	{AndExpression{BoolExpression(true), BoolExpression(true)}, "true && true", true},
	{AndExpression{BoolExpression(false), BoolExpression(true)}, "false && true", true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, "-1 && 1", false},

	{AndExpression{BoolExpression(true), BoolExpression(true)}, "false&&true", false},
	{AndExpression{BoolExpression(false), BoolExpression(true)}, "false&&false", false},
}

func TestAndPretty(t *testing.T) {
	for _, test := range testAndPrettyTests {
		if got := test.input.Pretty(); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestAndEval tests the Eval function

type TestAndEvalCase struct {
	input     AndExpression
	want      Value
	compliant bool
}

var testAndEvalTests = []TestAndEvalCase{
	{AndExpression{BoolExpression(true), BoolExpression(true)}, BoolValue(true), true},
	{AndExpression{BoolExpression(true), BoolExpression(false)}, BoolValue(false), true},
	{AndExpression{BoolExpression(false), BoolExpression(true)}, BoolValue(false), true},
	{AndExpression{BoolExpression(true), NumberExpression(1)}, UndefinedValue(), true},

	{AndExpression{NumberExpression(-1), BoolExpression(true)}, UndefinedValue(), true},
	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, UndefinedValue(), true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, IntValue(-1), false},
	{AndExpression{NumberExpression(1), NumberExpression(1)}, BoolValue(true), false},

	{AndExpression{NumberExpression(1), BoolExpression(false)}, IntValue(-1), false},
	{AndExpression{NumberExpression(1), BoolExpression(false)}, BoolValue(false), false},
}

func TestAndEval(t *testing.T) {
	for _, test := range testAndEvalTests {
		if got := test.input.Eval(&ValueState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t reflect %t", StructToJson(got), StructToJson(test.want), test.compliant, reflect.DeepEqual(got, test.want))
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
	{AndExpression{BoolExpression(false), BoolExpression(true)}, TypeBool, true},

	{AndExpression{NumberExpression(-1), BoolExpression(true)}, TypeIllTyped, true},
	{AndExpression{BoolExpression(true), NumberExpression(-1)}, TypeIllTyped, true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, TypeIllTyped, true},

	{AndExpression{NumberExpression(-1), NumberExpression(-1)}, TypeInt, false},
	{AndExpression{NumberExpression(0), NumberExpression(0)}, TypeInt, false},
}

// TestAndInfer tests the Infer function

func TestAndInfer(t *testing.T) {
	for _, test := range testAndInferTests {
		if got := test.input.Infer(&TypeState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
