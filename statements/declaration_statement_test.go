package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestDeclaration tests the Declaration function

type TestDeclarationCase struct {
	lhs       string
	rhs       Expression
	want      DeclarationStatement
	compliant bool
}

var testDeclarationTests = []TestDeclarationCase{

	{"x", NumberExpression(1), DeclarationStatement{"x", NumberExpression(1)}, true},
	{"x", NumberExpression(1), DeclarationStatement{"x", BoolExpression(true)}, false},
	{"x", PlusExpression{NumberExpression(1), BoolExpression(true)}, DeclarationStatement{"x", PlusExpression{NumberExpression(1), BoolExpression(true)}}, true},
}

func TestDeclaration(t *testing.T) {
	for _, test := range testDeclarationTests {
		got := Declaration(test.lhs, test.rhs)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestDeclarationEval tests the Eval function
type TestDeclarationEvalCase struct {
	input     DeclarationStatement
	want      ValueState
	compliant bool
}

var testDeclarationEvalTests = []TestDeclarationEvalCase{

	{DeclarationStatement{"x", NumberExpression(1)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, true},
	{DeclarationStatement{"x", BoolExpression(true)}, ValueState{map[string]Value{"x": {ValueBool, 0, true}}}, true},
	{DeclarationStatement{"x", PlusExpression{NumberExpression(1), BoolExpression(true)}}, ValueState{map[string]Value{"x": {Undefined, 0, false}}}, true},
	{DeclarationStatement{"x", PlusExpression{NumberExpression(1), BoolExpression(true)}}, ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, false},
	{DeclarationStatement{"x", NumberExpression(0)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, false},
	{DeclarationStatement{"x", BoolExpression(true)}, ValueState{map[string]Value{"x": {ValueInt, 0, true}}}, false},
}

func TestDeclarationEval(t *testing.T) {
	for _, test := range testDeclarationEvalTests {
		got := ValueState{map[string]Value{"x": {ValueInt, 1, false}}}
		test.input.Eval(&got)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestDeclarationPretty tests the Pretty function

type TestDeclarationPrettyCase struct {
	input     DeclarationStatement
	want      string
	compliant bool
}

var testDeclarationPrettyTests = []TestDeclarationPrettyCase{

	{DeclarationStatement{"x", NumberExpression(1)}, "x := 1", true},
	{DeclarationStatement{"x", PlusExpression{NumberExpression(1), BoolExpression(true)}}, "x := 1 + true", true},
	{DeclarationStatement{"x", BoolExpression(true)}, "x := 1", false},
}

func TestDeclarationPretty(t *testing.T) {
	for _, test := range testDeclarationPrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestDeclarationcheck tests the check function

type TestDeclarationCheckCase struct {
	input1    TypeState
	input2    DeclarationStatement
	want1     TypeState
	want2     bool
	compliant bool
}

var testDeclarationCheckTests = []TestDeclarationCheckCase{
	{TypeState{map[string]Type{"x": TypeInt}}, DeclarationStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt}}, true, true},
	{TypeState{map[string]Type{"x": TypeIllTyped}}, DeclarationStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt}}, true, true},
	{TypeState{map[string]Type{"x": TypeIllTyped}}, DeclarationStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeIllTyped}}, false, true},
	{TypeState{map[string]Type{"x": TypeInt}}, DeclarationStatement{"x", PlusExpression{NumberExpression(1), BoolExpression(true)}}, TypeState{map[string]Type{"x": TypeInt}}, true, true},
	{TypeState{map[string]Type{"x": TypeInt}}, DeclarationStatement{"y", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt, "y": TypeInt}}, true, true},
	{TypeState{map[string]Type{"x": TypeInt}}, DeclarationStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeBool}}, true, false},
	{TypeState{map[string]Type{"x": TypeInt}}, DeclarationStatement{"x", NumberExpression(1)}, TypeState{map[string]Type{"x": TypeInt}}, false, false},
}

func TestDeclarationCheck(t *testing.T) {
	for _, test := range testDeclarationCheckTests {
		got := test.input1
		if (test.input2.Check(&got) == test.want2) && reflect.DeepEqual(got, test.want1) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want1), test.compliant)
		}
	}
}
