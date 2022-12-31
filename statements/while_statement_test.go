package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestWhile tests the While function

type TestWhileCase struct {
	cond      Expression
	stmt      Statement
	want      WhileStatement
	compliant bool
}

var testWhileTests = []TestWhileCase{
	{BoolExpression(true), AssignmentStatement{"x", NumberExpression(1)}, WhileStatement{BoolExpression(true), AssignmentStatement{"x", NumberExpression(1)}}, true},
	{BoolExpression(true), AssignmentStatement{"x", NumberExpression(1)}, WhileStatement{BoolExpression(false), AssignmentStatement{"x", NumberExpression(1)}}, false},
	{BoolExpression(true), AssignmentStatement{"x", NumberExpression(1)}, WhileStatement{BoolExpression(true), AssignmentStatement{"x", NumberExpression(2)}}, false},
}

func TestWhile(t *testing.T) {
	for _, test := range testWhileTests {
		got := WhileStatement{test.cond, test.stmt}
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestWhileEval tests the Eval function

type TestWhileEvalCase struct {
	input     WhileStatement
	want      ValueState
	compliant bool
}

var testWhileEvalTests = []TestWhileEvalCase{
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, ValueState{"x": Value{ValueInt, 10, false}}, true},
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, ValueState{"x": Value{ValueInt, 100, false}}, false},
	{WhileStatement{NumberExpression(1), AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, ValueState{"x": Value{ValueInt, 100, false}}, false},
}

func TestEvalWhile(t *testing.T) {
	for _, test := range testWhileEvalTests {
		got := ValueState{"x": Value{ValueInt, 0, false}}
		test.input.Eval(got)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestWhilePretty tests the Pretty function

type TestWhilePrettyCase struct {
	input     WhileStatement
	want      string
	compliant bool
}

var testWhilePrettyTests = []TestWhilePrettyCase{
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, "while x < 10 { x = x + 1 }", true},
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, "while x < 10 { x = x + 2 }", false},
}

func TestWhilePretty(t *testing.T) {
	for _, test := range testWhilePrettyTests {
		vs := ValueState{"x": Value{ValueInt, 0, false}}
		if got := test.input.Pretty(vs); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestWhileCheck tests the Check function

type TestWhileCheckCase struct {
	input     WhileStatement
	want      TypeState
	compliant bool
}

var testWhileCheckTests = []TestWhileCheckCase{
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, TypeState{"x": TypeInt}, true},
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, TypeState{"x": TypeBool}, false},
	{WhileStatement{LesserExpression{Variable("y"), NumberExpression(10)}, AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}, TypeState{"x": TypeInt}, false},
}

func TestWhileCheck(t *testing.T) {
	for _, test := range testWhileCheckTests {
		got := TypeState{"x": TypeInt}
		if (reflect.DeepEqual(got, test.want) && test.input.Check(got)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}