package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestProgram tests the Program function

type TestProgramCase struct {
	input     Statement
	want      ProgramStatement
	compliant bool
}

var testProgramTests = []TestProgramCase{
	{DeclarationStatement{"x", NumberExpression(1)}, ProgramStatement{DeclarationStatement{"x", NumberExpression(1)}}, true},
	{DeclarationStatement{"x", NumberExpression(1)}, ProgramStatement{DeclarationStatement{"x", NumberExpression(2)}}, false},
}

func TestProgram(t *testing.T) {
	for _, test := range testProgramTests {
		got := ProgramStatement{test.input}
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestProgramEval tests the Eval function
type TestProgramEvalCase struct {
	input     ProgramStatement
	want      ValueState
	compliant bool
}

var testProgramEvalTests = []TestProgramEvalCase{
	{ProgramStatement{DeclarationStatement{"x", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 1, false}}, true},
	{ProgramStatement{DeclarationStatement{"x", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 2, false}}, false},
	{ProgramStatement{DeclarationStatement{"x", BoolExpression(true)}}, ValueState{"x": Value{ValueInt, 0, true}}, false},
}

func TestProgramEval(t *testing.T) {
	for _, test := range testProgramEvalTests {
		got := ValueState{"x": Value{ValueInt, 1, false}}

		test.input.Eval(got)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestProgramPretty tests the Pretty function

type TestProgramPrettyCase struct {
	input     ProgramStatement
	want      string
	compliant bool
}

var testProgramPrettyTests = []TestProgramPrettyCase{
	{ProgramStatement{DeclarationStatement{"x", BoolExpression(true)}}, "x := true", true},
	{ProgramStatement{DeclarationStatement{"x", BoolExpression(true)}}, "{x := true}", false},
}

func TestProgramPretty(t *testing.T) {
	for _, test := range testProgramPrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestProgramCheck tests the Check function

type TestProgramCheckCase struct {
	input1    TypeState
	input2    ProgramStatement
	want1     TypeState
	want2     bool
	compliant bool
}

var testProgramCheckTests = []TestProgramCheckCase{
	{TypeState{"x": TypeInt}, ProgramStatement{DeclarationStatement{"x", NumberExpression(1)}}, TypeState{"x": TypeInt}, true, true},
	{TypeState{"x": TypeInt}, ProgramStatement{DeclarationStatement{"x", NumberExpression(1)}}, TypeState{"x": TypeInt}, false, false},
	{TypeState{"x": TypeInt}, ProgramStatement{DeclarationStatement{"x", NumberExpression(1)}}, TypeState{"x": TypeBool}, true, false},
}

func TestProgramCheck(t *testing.T) {
	for _, test := range testProgramCheckTests {
		got := test.input1
		if (reflect.DeepEqual(got, test.want1) && test.input2.Check(got) == test.want2) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want1), test.compliant)
		}
	}
}
