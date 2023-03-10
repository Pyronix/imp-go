package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestBool tests the Bool function

type TestBoolCase struct {
	input     bool
	want      BoolExpression
	compliant bool
}

var testBoolTests = []TestBoolCase{
	{true, BoolExpression(true), true},
	{false, BoolExpression(false), true},
	{true, BoolExpression(false), false},
	{false, BoolExpression(true), false},
}

func TestBool(t *testing.T) {
	for _, test := range testBoolTests {
		if got := Bool(test.input); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestBoolPretty tests the Pretty function

type TestBoolPrettyCase struct {
	input     BoolExpression
	want      string
	compliant bool
}

var testBoolPrettyTests = []TestBoolPrettyCase{
	{BoolExpression(true), "true", true},
	{BoolExpression(false), "false", true},
	{BoolExpression(true), "false", false},
	{BoolExpression(false), "true", false},
}

func TestBoolPretty(t *testing.T) {
	for _, test := range testBoolPrettyTests {
		if got := test.input.Pretty(); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestBoolEval tests the Eval function

type TestBoolEvalCase struct {
	input     BoolExpression
	want      Value
	compliant bool
}

var testBoolEvalTests = []TestBoolEvalCase{
	{BoolExpression(true), BoolValue(true), true},
	{BoolExpression(false), BoolValue(false), true},
	{BoolExpression(true), BoolValue(false), false},
	{BoolExpression(false), BoolValue(true), false},
}

func TestBoolEval(t *testing.T) {
	for _, test := range testBoolEvalTests {
		if got := test.input.Eval(&ValueState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

type TestBoolInferCase struct {
	input     BoolExpression
	want      Type
	compliant bool
}

var testBoolInferTests = []TestBoolInferCase{
	{BoolExpression(false), TypeBool, true},
	{BoolExpression(true), TypeBool, true},
	{BoolExpression(true), TypeInt, false},
	{BoolExpression(false), TypeInt, false},
}

// TestBoolInfer tests the Infer function

func TestBoolInfer(t *testing.T) {
	for _, test := range testBoolInferTests {
		if got := test.input.Infer(&TypeState{}); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
