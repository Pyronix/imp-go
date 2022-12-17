package types

import (
	. "imp/helper"
	"testing"
)

// TestNumber tests the Number function

type TestNumberCase struct {
	arg1     int
	expected NumberExpression
}

var testNumberTests = []TestNumberCase{
	TestNumberCase{-1, NumberExpression(-1)},
	TestNumberCase{0, NumberExpression(0)},
	TestNumberCase{1, NumberExpression(1)},
}

func TestNumber(t *testing.T) {
	for _, test := range testNumberTests {
		if output := Number(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

// TestPretty tests the Pretty function

type TestPrettyCase struct {
	arg1     NumberExpression
	expected string
}

var CompliantTestPretty = []TestPrettyCase{
	TestPrettyCase{NumberExpression(-1), "-1"},
	TestPrettyCase{NumberExpression(0), "0"},
	TestPrettyCase{NumberExpression(1), "1"},
}

var NonCompliantTestPretty = []TestPrettyCase{
	TestPrettyCase{NumberExpression(-1), "-10"},
	TestPrettyCase{NumberExpression(0), "999"},
	TestPrettyCase{NumberExpression(1), "10"},
}

func TestPretty(t *testing.T) {
	for _, test := range CompliantTestPretty {
		if output := test.arg1.Pretty(); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}

	for _, test := range NonCompliantTestPretty {
		if output := test.arg1.Pretty(); output == test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

// TestEval tests the Eval function

type TestEvalCase struct {
	arg1     NumberExpression
	expected Value
}

var CompliantTestEval = []TestEvalCase{
	TestEvalCase{NumberExpression(-1), IntValue(-1)},
	TestEvalCase{NumberExpression(0), IntValue(0)},
	TestEvalCase{NumberExpression(1), IntValue(1)},
}

var NonCompliantTestEval = []TestEvalCase{
	TestEvalCase{NumberExpression(-1), IntValue(-10)},
	TestEvalCase{NumberExpression(0), IntValue(999)},
	TestEvalCase{NumberExpression(1), IntValue(10)},
}

func TestEval(t *testing.T) {
	for _, test := range CompliantTestEval {
		if output := test.arg1.Eval(ValueState{}); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", StructToJson(output), StructToJson(test.expected))
		}
	}

	for _, test := range NonCompliantTestEval {
		if output := test.arg1.Eval(ValueState{}); output == test.expected {
			t.Errorf("Output %q not equal to expected %q", StructToJson(output), StructToJson(test.expected))
		}
	}
}

type TestInferCase struct {
	arg1     NumberExpression
	expected Type
}

var CompliantTestInfer = []TestInferCase{
	TestInferCase{NumberExpression(-1), TypeInt},
	TestInferCase{NumberExpression(0), TypeInt},
	TestInferCase{NumberExpression(1), TypeInt},
}

var NonCompliantTestInfer = []TestInferCase{
	TestInferCase{NumberExpression(-1), TypeBool},
	TestInferCase{NumberExpression(0), TypeBool},
	TestInferCase{NumberExpression(1), TypeBool},
}

// TestInfer tests the Infer function

func TestInfer(t *testing.T) {
	for _, test := range CompliantTestInfer {
		if output := test.arg1.Infer(TypeState{}); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
	for _, test := range NonCompliantTestInfer {
		if output := test.arg1.Infer(TypeState{}); output == test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
