package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestSequence tests the While function

type TestSequenceCase struct {
	stmt1     Statement
	stmt2     Statement
	want      SequenceStatement
	compliant bool
}

var testSequenceTests = []TestSequenceCase{
	{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}}, true},
	{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(2)}}, false},
	{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}, SequenceStatement{AssignmentStatement{"x", NumberExpression(2)}, AssignmentStatement{"y", NumberExpression(1)}}, false},
	{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}, SequenceStatement{AssignmentStatement{"x", NumberExpression(2)}, AssignmentStatement{"y", NumberExpression(2)}}, false},
}

func TestSequence(t *testing.T) {
	for _, test := range testSequenceTests {
		got := SequenceStatement{test.stmt1, test.stmt2}
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestSequenceEval tests the Eval function

type TestSequenceEvalCase struct {
	input     ValueState
	sequence  SequenceStatement
	want      ValueState
	compliant bool
}

var testSequenceEvalTests = []TestSequenceEvalCase{
	{ValueState{"x": Value{ValueInt, 0, false}}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 1, false}, "y": Value{ValueInt, 1, false}}, true},
	{ValueState{"x": Value{ValueInt, 0, false}}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 1, false}, "y": Value{ValueInt, 2, false}}, false},
	{ValueState{"x": Value{ValueInt, 0, false}}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 2, false}, "y": Value{ValueInt, 1, false}}, false},
	{ValueState{"x": Value{ValueInt, 0, false}}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 2, false}, "y": Value{ValueInt, 2, false}}, false},
}

func TestSequenceEval(t *testing.T) {
	for _, test := range testSequenceEvalTests {
		test.sequence.Eval(test.input)
		if reflect.DeepEqual(test.input, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(test.input), StructToJson(test.want), test.compliant)
		}
	}
}

// TestSequencePretty tests the Pretty function

type TestSequencePrettyCase struct {
	sequence  SequenceStatement
	want      string
	compliant bool
}

var testSequencePrettyTests = []TestSequencePrettyCase{
	{SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}}, "x = 1; y = 1", true},
	{SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}}, "x = 1; y = 2", false},
	{SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}}, "x = 2; y = 1", false},
	{SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, AssignmentStatement{"y", NumberExpression(1)}}, "x = 2; y = 2", false},
}

func TestSequencePretty(t *testing.T) {
	for _, test := range testSequencePrettyTests {
		if got := test.sequence.Pretty(); got == test.want != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestSequenceCheck tests the Check function

type TestSequenceCheckCase struct {
	input     TypeState
	sequence  SequenceStatement
	want1     TypeState
	want2     bool
	compliant bool
}

var testSequenceCheckTests = []TestSequenceCheckCase{
	{TypeState{"x": TypeInt}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt, "y": TypeInt}, true, true},
	{TypeState{}, SequenceStatement{DeclarationStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt, "y": TypeInt}, true, true},
	{TypeState{"x": TypeInt}, SequenceStatement{AssignmentStatement{"x", BoolExpression(true)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt}, false, true},
	{TypeState{"x": TypeInt}, SequenceStatement{AssignmentStatement{"x", BoolExpression(true)}, DeclarationStatement{"y", BoolExpression(true)}}, TypeState{"x": TypeInt}, false, true},
	{TypeState{"x": TypeInt}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt, "y": TypeInt}, false, false},
	{TypeState{}, SequenceStatement{DeclarationStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt, "y": TypeInt}, false, false},
	{TypeState{"x": TypeInt}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt, "y": TypeInt}, false, false},
	{TypeState{"x": TypeInt}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt, "y": TypeBool}, true, false},
	{TypeState{"x": TypeInt}, SequenceStatement{AssignmentStatement{"x", NumberExpression(1)}, DeclarationStatement{"y", NumberExpression(1)}}, TypeState{"x": TypeInt, "y": TypeBool}, false, false},
}

func TestSequenceCheck(t *testing.T) {
	for _, test := range testSequenceCheckTests {
		if ((test.sequence.Check(test.input) == test.want2) && reflect.DeepEqual(test.input, test.want1)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(test.input), StructToJson(test.want1), test.compliant)
		}
	}
}
