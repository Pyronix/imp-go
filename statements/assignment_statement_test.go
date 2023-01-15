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
		got := Assignment(test.lhs, test.rhs)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestAssigmentEval tests the Eval function
type TestAssigmentEvalCase struct {
	input     AssignmentStatement
	want      ValueState
	panic     bool
	compliant bool
}

var testAssignmentEvalTests = []TestAssigmentEvalCase{

	{AssignmentStatement{"x", NumberExpression(1)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, false, true},
	{AssignmentStatement{"x", NumberExpression(0)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, false, false},
	{AssignmentStatement{"x", BoolExpression(true)}, ValueState{map[string]Value{"x": {ValueInt, 0, true}}}, true, false},
}

func TestAssignmentEval(t *testing.T) {
	for _, test := range testAssignmentEvalTests {
		got := ValueState{map[string]Value{"x": {ValueInt, 1, false}}}
		if test.panic {
			func() {
				defer func() { _ = recover() }()
				test.input.Eval(&got)
				t.Errorf("expected panic but it is not")
			}()
		} else {
			test.input.Eval(&got)
			if reflect.DeepEqual(got, test.want) != test.compliant {
				t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
			}
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
	{TypeState{map[string]Type{"x": TypeInt}}, AssignmentStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt}}, true, true},
	{TypeState{map[string]Type{"x": TypeIllTyped}}, AssignmentStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt}}, false, false},
	{TypeState{map[string]Type{"x": TypeIllTyped}}, AssignmentStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt}}, true, false},
	{TypeState{map[string]Type{"x": TypeInt}}, AssignmentStatement{"x1", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt}}, true, false},

	{TypeState{map[string]Type{"x": TypeIllTyped}}, AssignmentStatement{"x", EqualityExpression{NumberExpression(1), BoolExpression(false)}}, TypeState{map[string]Type{"x": TypeIllTyped}}, false, true},
	{TypeState{map[string]Type{"x": TypeInt}}, AssignmentStatement{"x", EqualityExpression{NumberExpression(1), BoolExpression(false)}}, TypeState{map[string]Type{"x": TypeIllTyped}}, false, false},
}

func TestAssignmentCheck(t *testing.T) {
	for _, test := range testAssignmentCheckTests {
		got := test.input1
		if (reflect.DeepEqual(got, test.want1) && test.input2.Check(&got) == test.want2) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want1), test.compliant)
		}
	}
}
