package types

import (
	. "imp/helper"
	"testing"
)

// TestEqual tests the Equal function

type TestGroupingCase struct {
	input     Expression
	want      GroupingExpression
	compliant bool
}

var testGroupingTests = []TestGroupingCase{
	{BoolExpression(true), GroupingExpression{BoolExpression(true)}, true},
	{NumberExpression(1), GroupingExpression{NumberExpression(1)}, true},
	{AndExpression{BoolExpression(true), BoolExpression(true)}, GroupingExpression{AndExpression{BoolExpression(true), BoolExpression(true)}}, false},

	{BoolExpression(true), GroupingExpression{NumberExpression(1)}, false},
	{NumberExpression(1), GroupingExpression{BoolExpression(true)}, false},
}

func TestGrouping(t *testing.T) {
	for _, test := range testGroupingTests {
		if got := Grouping(test.input); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestGroupingPretty tests the Pretty function

type TestGroupingPrettyCase struct {
	input     GroupingExpression
	want      string
	compliant bool
}

var testGroupingPrettyTests = []TestGroupingPrettyCase{
	{GroupingExpression{NumberExpression(-1)}, "(-1)", true},
	{GroupingExpression{OrExpression{BoolExpression(true), BoolExpression(true)}}, "(true||true)", true},

	{GroupingExpression{OrExpression{BoolExpression(true), BoolExpression(true)}}, "???", false},
}

func TestGroupingPretty(t *testing.T) {
	for _, test := range testGroupingPrettyTests {
		if got := test.input.Pretty(); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}

// TestGroupingEval tests the Eval function

type TestGroupingEvalCase struct {
	input     GroupingExpression
	want      Value
	compliant bool
}

var testGroupingEvalTests = []TestGroupingEvalCase{
	{GroupingExpression{NumberExpression(-1)}, IntValue(-1), true},
	{GroupingExpression{OrExpression{BoolExpression(true), BoolExpression(true)}}, BoolValue(true), true},
	{GroupingExpression{OrExpression{BoolExpression(false), BoolExpression(true)}}, BoolValue(true), true},
	{GroupingExpression{OrExpression{BoolExpression(true), BoolExpression(false)}}, BoolValue(true), true},

	{GroupingExpression{PlusExpression{BoolExpression(true), BoolExpression(true)}}, UndefinedValue(), true},
	{GroupingExpression{PlusExpression{NumberExpression(-1), NumberExpression(-1)}}, IntValue(-2), true},
	{GroupingExpression{PlusExpression{NumberExpression(-1), BoolExpression(true)}}, UndefinedValue(), true},
	{GroupingExpression{PlusExpression{BoolExpression(true), NumberExpression(-1)}}, UndefinedValue(), true},

	{GroupingExpression{OrExpression{BoolExpression(true), BoolExpression(true)}}, UndefinedValue(), false},
}

func TestGroupingEval(t *testing.T) {
	for _, test := range testGroupingEvalTests {
		if got := test.input.Eval(ValueState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", StructToJson(got), StructToJson(test.want))
		}
	}
}

type TestGroupingInferCase struct {
	input     GroupingExpression
	want      Type
	compliant bool
}

var testGroupingInferTests = []TestGroupingInferCase{
	{GroupingExpression{NumberExpression(-1)}, TypeInt, true},
	{GroupingExpression{OrExpression{BoolExpression(true), BoolExpression(true)}}, TypeBool, true},
	{GroupingExpression{OrExpression{NumberExpression(-1), BoolExpression(true)}}, TypeIllTyped, true},

	{GroupingExpression{NumberExpression(-1)}, TypeIllTyped, false},
	{GroupingExpression{OrExpression{BoolExpression(true), BoolExpression(true)}}, TypeInt, false},
	{GroupingExpression{OrExpression{NumberExpression(-1), BoolExpression(true)}}, TypeBool, false},
}

// TestGroupingInfer tests the Infer function

func TestGroupingInfer(t *testing.T) {
	for _, test := range testGroupingInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
