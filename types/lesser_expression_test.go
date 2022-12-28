package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestLesser tests the Lesser function

type TestLesserCase struct {
	input     Expression
	input2    Expression
	want      LesserExpression
	compliant bool
}

var testLesserTests = []TestLesserCase{
	{NumberExpression(1), NumberExpression(2), LesserExpression{NumberExpression(1), NumberExpression(2)}, true},
	{NumberExpression(1), NumberExpression(2), LesserExpression{NumberExpression(2), NumberExpression(1)}, false},
}

func TestLesser(t *testing.T) {
	for _, test := range testLesserTests {
		if got := Lesser(test.input, test.input2); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestLesserPretty tests the Pretty function

type TestLesserPrettyCase struct {
	input     LesserExpression
	want      string
	compliant bool
}

var testLesserPrettyTests = []TestLesserPrettyCase{
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, "1 < 2", true},
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, "2 < 1", false},
}

func TestLesserPretty(t *testing.T) {
	for _, test := range testLesserPrettyTests {
		if got := test.input.Pretty(); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestLesserEval tests the Eval function

type TestLesserEvalCase struct {
	input     Expression
	want      Value
	compliant bool
}

var TestLesserEvalTests = []TestLesserEvalCase{
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, BoolValue(true), true},
	{LesserExpression{NumberExpression(1), BoolExpression(true)}, UndefinedValue(), true},
	{LesserExpression{NumberExpression(2), NumberExpression(1)}, BoolValue(false), true},
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, BoolValue(false), false},
}

func TestLesserEval(t *testing.T) {
	for _, test := range TestLesserEvalTests {
		if got := test.input.Eval(ValueState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

type TestLesserInferCase struct {
	input     LesserExpression
	want      Type
	compliant bool
}

var TestLesserInferTests = []TestLesserInferCase{
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, TypeBool, true},
	{LesserExpression{NumberExpression(1), BoolExpression(true)}, TypeIllTyped, true},
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, TypeInt, false},
}

// TestLesserInfer tests the Infer function

func TestLesserInfer(t *testing.T) {
	for _, test := range TestLesserInferTests {
		if got := test.input.Infer(TypeState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %q not equal to want %q, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
