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
	TestNumberCase{-1, NumberExpression(-1), true},
	TestNumberCase{0, NumberExpression(0), true},
	TestNumberCase{1, NumberExpression(1), true},
	TestNumberCase{-1, NumberExpression(-10), false},
	TestNumberCase{0, NumberExpression(999), false},
	TestNumberCase{1, NumberExpression(10), false},
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
	TestPrettyCase{NumberExpression(-1), "-1", true},
	TestPrettyCase{NumberExpression(0), "0", true},
	TestPrettyCase{NumberExpression(1), "1", true},
	TestPrettyCase{NumberExpression(-1), "-10", false},
	TestPrettyCase{NumberExpression(0), "999", false},
	TestPrettyCase{NumberExpression(1), "10", false},
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
	TestEvalCase{NumberExpression(-1), IntValue(-1), true},
	TestEvalCase{NumberExpression(0), IntValue(0), true},
	TestEvalCase{NumberExpression(1), IntValue(1), true},
	TestEvalCase{NumberExpression(-1), IntValue(-10), false},
	TestEvalCase{NumberExpression(0), IntValue(999), false},
	TestEvalCase{NumberExpression(1), IntValue(10), false},
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
	TestInferCase{NumberExpression(-1), TypeInt, true},
	TestInferCase{NumberExpression(0), TypeInt, true},
	TestInferCase{NumberExpression(1), TypeInt, true},
	TestInferCase{NumberExpression(-1), TypeBool, false},
	TestInferCase{NumberExpression(0), TypeBool, false},
	TestInferCase{NumberExpression(1), TypeBool, false},
}

// TestInfer tests the Infer function

func TestInfer(t *testing.T) {
	for _, test := range testInferTests {
		if got := test.input.Infer(TypeState{}); got != test.want && test.compliant {
			t.Errorf("got %q not equal to want %q", got, test.want)
		}
	}
}
