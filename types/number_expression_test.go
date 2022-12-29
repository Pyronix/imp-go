package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestNumber tests the Bool function

type TestNumberCase struct {
	input     int
	want      Expression
	compliant bool
}

var testNumberTests = []TestNumberCase{
	{1, NumberExpression(1), true},

	{0, BoolExpression(true), false},
	{0, NumberExpression(1), false},
}

func TestNumber(t *testing.T) {
	for _, test := range testNumberTests {
		if got := Number(test.input); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestNumberPretty tests the Pretty function

type TestNumberPrettyCase struct {
	input     NumberExpression
	want      string
	compliant bool
}

var testNumberPrettyTests = []TestNumberPrettyCase{
	{NumberExpression(-1), "-1", true},
	{NumberExpression(0), "0", true},
	{NumberExpression(1), "1", true},
	{NumberExpression(-1), "-10", false},
	{NumberExpression(0), "999", false},
	{NumberExpression(1), "10", false},
}

func TestNumberPretty(t *testing.T) {
	for _, test := range testNumberPrettyTests {
		if got := test.input.Pretty(); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestNumberEval tests the Eval function

type TestNumberEvalCase struct {
	input     NumberExpression
	want      Value
	compliant bool
}

var testNumberEvalTests = []TestNumberEvalCase{
	{NumberExpression(-1), IntValue(-1), true},
	{NumberExpression(0), IntValue(0), true},
	{NumberExpression(1), IntValue(1), true},
	{NumberExpression(-1), IntValue(-10), false},
	{NumberExpression(0), IntValue(999), false},
	{NumberExpression(1), IntValue(10), false},
}

func TestEval(t *testing.T) {
	for _, test := range testNumberEvalTests {
		if got := test.input.Eval(ValueState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

type TestNumberInferCase struct {
	input     NumberExpression
	want      Type
	compliant bool
}

var testNumberInferTests = []TestNumberInferCase{
	{NumberExpression(-1), TypeInt, true},
	{NumberExpression(0), TypeInt, true},
	{NumberExpression(1), TypeInt, true},
	{NumberExpression(-1), TypeBool, false},
	{NumberExpression(0), TypeBool, false},
	{NumberExpression(1), TypeBool, false},
}

// TestNumberInfer tests the Infer function

func TestNumberInfer(t *testing.T) {
	for _, test := range testNumberInferTests {
		if got := test.input.Infer(TypeState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
