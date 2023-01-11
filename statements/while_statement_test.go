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
	bockStmt  BlockStatement
	want      WhileStatement
	compliant bool
}

var testWhileTests = []TestWhileCase{
	{BoolExpression(true), BlockStatement{AssignmentStatement{"x", NumberExpression(1)}}, WhileStatement{BoolExpression(true), BlockStatement{AssignmentStatement{"x", NumberExpression(1)}}}, true},
	{BoolExpression(true), BlockStatement{AssignmentStatement{"x", NumberExpression(1)}}, WhileStatement{BoolExpression(false), BlockStatement{AssignmentStatement{"x", NumberExpression(1)}}}, false},
	{BoolExpression(true), BlockStatement{AssignmentStatement{"x", NumberExpression(1)}}, WhileStatement{BoolExpression(true), BlockStatement{AssignmentStatement{"x", NumberExpression(2)}}}, false},
}

func TestWhile(t *testing.T) {
	for _, test := range testWhileTests {
		got := While(test.cond, test.bockStmt)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestWhileEval tests the Eval function

type TestWhileEvalCase struct {
	input     WhileStatement
	want      ValueState
	panic     bool
	compliant bool
}

var testWhileEvalTests = []TestWhileEvalCase{
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, ValueState{map[string]Value{"x": {ValueInt, 10, false}}}, false, true},
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, ValueState{map[string]Value{"x": {ValueInt, 100, false}}}, false, false},
	{WhileStatement{NumberExpression(1), BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, ValueState{map[string]Value{"x": {ValueInt, 100, false}}}, true, false},
}

func TestEvalWhile(t *testing.T) {
	for _, test := range testWhileEvalTests {
		got := ValueState{map[string]Value{"x": {ValueInt, 0, false}}}
		func() {
			defer func() { _ = recover() }()
			test.input.Eval(&got)
			if reflect.DeepEqual(got, test.want) != test.compliant || test.panic != (recover() != nil) {
				t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
			}
		}()

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
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, "while x < 10 {x = x + 1}", true},
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, "while x < 10 {x = x + 2}", false},
}

func TestWhilePretty(t *testing.T) {
	for _, test := range testWhilePrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
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
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, TypeState{map[string]Type{"x": TypeInt}}, true},
	{WhileStatement{LesserExpression{Variable("x"), NumberExpression(10)}, BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, TypeState{map[string]Type{"x": TypeBool}}, false},
	{WhileStatement{LesserExpression{Variable("y"), NumberExpression(10)}, BlockStatement{AssignmentStatement{"x", PlusExpression{Variable("x"), NumberExpression(1)}}}}, TypeState{map[string]Type{"x": TypeInt}}, false},
}

func TestWhileCheck(t *testing.T) {
	for _, test := range testWhileCheckTests {
		got := TypeState{map[string]Type{"x": TypeInt}}
		if (reflect.DeepEqual(got, test.want) && test.input.Check(&got)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}
