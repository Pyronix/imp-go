package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestAssignment tests the Assigment function

type TestAssignmentCase struct {
	lhs       string
	rhs       Expression
	want      AssignmentStatement
	compliant bool
}

var testAssignmentTests = []TestAssignmentCase{

	{"temp1", NumberExpression(1), AssignmentStatement{"temp1", NumberExpression(1)}, true},
	{"temp1", NumberExpression(1), AssignmentStatement{"temp1", BoolExpression(true)}, false},
}

func TestAssignment(t *testing.T) {
	for _, test := range testAssignmentTests {
		got := AssignmentStatement{test.lhs, test.rhs}
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestAssigmentEval tests the Eval function
type TestAssigmentEvalCase struct {
	input     AssignmentStatement
	want      ValueState
	compliant bool
}

var testAssignmentEvalTests = []TestAssigmentEvalCase{

	{AssignmentStatement{"temp", NumberExpression(1)}, ValueState{"temp": Value{ValueInt, 1, false}}, true},
	{AssignmentStatement{"temp", NumberExpression(0)}, ValueState{"temp": Value{ValueInt, 1, false}}, false},
	{AssignmentStatement{"temp", BoolExpression(true)}, ValueState{"temp": Value{ValueInt, 0, true}}, false},
}

// Wie einen Fall mit valuestate {"temp": Value{undefined, 0, false}}
func TestAssignmentEval(t *testing.T) {
	for _, test := range testAssignmentEvalTests {
		got := ValueState{"temp": Value{ValueInt, 1, false}}
		//eval gibt nichts zurück
		//aber printet "Assignment Eval fail" wenns nicht klappt
		test.input.Eval(got)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestAssignmentPretty tests the Pretty function

type TestAssignmentPrettyCase struct {
	input     AssignmentStatement
	want      string
	compliant bool
}

var testAssignmentPrettyTests = []TestAssignmentPrettyCase{

	{AssignmentStatement{"temp", NumberExpression(1)}, "temp = 1", true},
	{AssignmentStatement{"temp", BoolExpression(true)}, "temp = 1", false},
}

func TestAssignmentPretty(t *testing.T) {
	for _, test := range testAssignmentPrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestAssignmentInfer tests the Infer function

type TestAssignmentInferCase struct {
	input2    AssignmentStatement
	want      TypeState
	compliant bool
}

var testAssignmentInferTests = []TestAssignmentInferCase{
	{AssignmentStatement{"temp", NumberExpression(1)}, TypeState{"temp": TypeInt}, true},
	{AssignmentStatement{"temp", NumberExpression(1)}, TypeState{"temp": TypeIllTyped}, false},
	{AssignmentStatement{"temp", NumberExpression(1)}, TypeState{"temp": TypeBool}, false},
}

func TestAssignmentInfer(t *testing.T) {
	for _, test := range testAssignmentInferTests {
		got := TypeState{"temp": TypeInt}
		if (reflect.DeepEqual(got, test.want) && test.input2.Check(got)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
