package types

import (
	"imp/helper"
	"testing"
)

// TestVariablePretty tests the Pretty function

type TestVariablePrettyCase struct {
	input     Variable
	want      string
	compliant bool
}

var testVariablePrettyTests = []TestVariablePrettyCase{
	{Variable("x"), "x", true},
	{Variable("y1"), "y1", true},
	{Variable("z_"), "z_", true},
	{Variable("z_"), "???", false},
	{Variable("z"), "???", false},
}

func TestVariablePretty(t *testing.T) {
	for _, test := range testVariablePrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestVariableEval tests the Variable function

type TestVariableEvalCase struct {
	input     Variable
	want      Value
	vs        ValueState
	compliant bool
}

var vs = ValueState{
	"x": Value{ValueType: ValueInt, IntValue: 1},
	"y": Value{ValueType: ValueBool, BoolValue: true},
	"z": Value{ValueType: Undefined},
}

var testVariableEvalTests = []TestVariableEvalCase{
	{Variable("x"), Value{ValueType: ValueInt, IntValue: 1}, vs, true},
	{Variable("y"), Value{ValueType: ValueBool, BoolValue: true}, vs, true},
	{Variable("z"), Value{ValueType: Undefined}, vs, true},
	{Variable("k"), Value{ValueType: Undefined}, vs, true},
	{Variable("z"), Value{ValueType: ValueInt, IntValue: 1}, vs, false},
	{Variable("z"), Value{ValueType: ValueBool, BoolValue: true}, vs, false},
}

func TestVariableEval(t *testing.T) {
	for _, test := range testVariableEvalTests {
		if got := test.input.Eval(test.vs); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", helper.StructToJson(got), helper.StructToJson(test.want))
		}
	}
}

// TestVariableInfer tests the Variable function

type TestVariableInferCase struct {
	input     Variable
	want      Type
	vs        TypeState
	compliant bool
}

var ts = TypeState{
	"x": TypeInt,
	"y": TypeBool,
	"z": TypeIllTyped,
}

var testVariableInferTests = []TestVariableInferCase{
	{Variable("x"), TypeInt, ts, true},
	{Variable("y"), TypeBool, ts, true},
	{Variable("z"), TypeIllTyped, ts, true},
	{Variable("k"), TypeIllTyped, ts, true},
	{Variable("x"), TypeBool, ts, false},
	{Variable("y"), TypeInt, ts, false},
	{Variable("z"), TypeInt, ts, false},
}

func TestVariableInfer(t *testing.T) {
	for _, test := range testVariableInferTests {
		if got := test.input.Infer(test.vs); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", helper.StructToJson(got), helper.StructToJson(test.want))
		}
	}
}
