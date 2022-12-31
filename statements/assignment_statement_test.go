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

	{"x", NumberExpression(1), AssignmentStatement{"x", NumberExpression(1)}, true},
	{"x", NumberExpression(1), AssignmentStatement{"x", BoolExpression(true)}, false},
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

	{AssignmentStatement{"x", NumberExpression(1)}, ValueState{"x": Value{ValueInt, 1, false}}, true},
	{AssignmentStatement{"x", NumberExpression(0)}, ValueState{"x": Value{ValueInt, 1, false}}, false},
	{AssignmentStatement{"x", BoolExpression(true)}, ValueState{"x": Value{ValueInt, 0, true}}, false},
}

func TestAssignmentEval(t *testing.T) {
	for _, test := range testAssignmentEvalTests {
		got := ValueState{"x": Value{ValueInt, 1, false}}
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

	{AssignmentStatement{"x", NumberExpression(1)}, "x = 1", true},
	{AssignmentStatement{"x", BoolExpression(true)}, "x = 1", false},
}

func TestAssignmentPretty(t *testing.T) {
	for _, test := range testAssignmentPrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestAssignmentCheck tests the Check function

type TestAssignmentCheckCase struct {
	input1    TypeState
	input2    AssignmentStatement
	want1     TypeState
	want2     bool
	compliant bool
}

var testAssignmentCheckTests = []TestAssignmentCheckCase{
	{TypeState{"x": TypeInt}, AssignmentStatement{"x", NumberExpression(1)}, TypeState{"x": TypeInt}, true, true},
	{TypeState{"x": TypeIllTyped}, AssignmentStatement{"x", NumberExpression(1)}, TypeState{"x": TypeInt}, false, false},
	{TypeState{"x": TypeIllTyped}, AssignmentStatement{"x", NumberExpression(1)}, TypeState{"x": TypeInt}, true, false},
	{TypeState{"x": TypeInt}, AssignmentStatement{"x1", NumberExpression(1)}, TypeState{"x": TypeInt}, true, false},

	{TypeState{"x": TypeIllTyped}, AssignmentStatement{"x", EqualityExpression{NumberExpression(1), BoolExpression(false)}}, TypeState{"x": TypeIllTyped}, false, true},
	{TypeState{"x": TypeInt}, AssignmentStatement{"x", EqualityExpression{NumberExpression(1), BoolExpression(false)}}, TypeState{"x": TypeIllTyped}, false, false},
}

func TestAssignmentCheck(t *testing.T) {
	for _, test := range testAssignmentCheckTests {
		got := test.input1
		if (reflect.DeepEqual(got, test.want1) && test.input2.Check(got) == test.want2) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want1), test.compliant)
		}
	}
}
