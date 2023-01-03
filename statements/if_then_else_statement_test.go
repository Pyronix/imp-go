package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestIfThenElse tests the IfThenElse function

type TestIfThenElseCase struct {
	cond      Expression
	thenStmt  BlockStatement
	elseStmt  BlockStatement
	want      IfThenElseStatement
	compliant bool
}

var testIfThenElseTests = []TestIfThenElseCase{
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}, IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, true},
	{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}, IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(2)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, false},
}

func TestIfThenElse(t *testing.T) {
	for _, test := range testIfThenElseTests {
		got := Ite(test.cond, test.thenStmt, test.elseStmt)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestIfThenElseEval tests the Eval function

type TestIfThenElseEvalCase struct {
	input     IfThenElseStatement
	vs        ValueState
	want      ValueState
	compliant bool
}

var testIfThenElseEvalTests = []TestIfThenElseEvalCase{
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, true},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 2, false}}}, false},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), BoolExpression(true)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 2, false}}}, false},
	{IfThenElseStatement{LesserExpression{NumberExpression(2), NumberExpression(1)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, ValueState{map[string]Value{"x": {ValueInt, 0, false}, "y": {ValueInt, 2, false}}}, true},
	{IfThenElseStatement{LesserExpression{NumberExpression(2), NumberExpression(1)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 0, false}, "y": {ValueInt, 2, false}}}, false},
}

func TestEvalIfThenElse(t *testing.T) {
	for _, test := range testIfThenElseEvalTests {
		test.input.Eval(&test.vs)
		if reflect.DeepEqual(test.vs, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(test.vs), StructToJson(test.want), test.compliant)
		}
	}
}

// TestIfThenElsePretty tests the Pretty function

type TestIfThenElsePrettyCase struct {
	input     IfThenElseStatement
	want      string
	compliant bool
}

var testIfThenElsePrettyTests = []TestIfThenElsePrettyCase{
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, "if 1 < 2 {x := 1} else {y := 2}", true},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, "if 1 < 2 {x := 1} else {y := 2}", true},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, "if 1 < 2 {x := 1} else {y := 2} ", false},
}

func TestIfThenElsePretty(t *testing.T) {
	for _, test := range testIfThenElsePrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestIfThenElseCheck tests the Check function

type TestIfThenElseCheckCase struct {
	input     IfThenElseStatement
	ts        TypeState
	want1     TypeState
	want2     bool
	compliant bool
}

var testIfThenElseCheckTests = []TestIfThenElseCheckCase{
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, TypeState{map[string]Type{"x": TypeBool, "y": TypeBool}}, TypeState{map[string]Type{"x": TypeInt, "y": TypeBool}}, true, true},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, TypeState{map[string]Type{"x": TypeInt, "y": TypeInt}}, TypeState{map[string]Type{"x": TypeInt, "y": TypeInt}}, true, true},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, TypeState{map[string]Type{"x": TypeInt, "y": TypeBool}}, TypeState{map[string]Type{"x": TypeInt, "y": TypeBool}}, true, false},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), NumberExpression(2)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, TypeState{map[string]Type{"x": TypeBool, "y": TypeInt}}, TypeState{map[string]Type{"x": TypeBool, "y": TypeInt}}, true, false},
	{IfThenElseStatement{LesserExpression{NumberExpression(1), BoolExpression(true)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, BlockStatement{DeclarationStatement{"y", NumberExpression(2)}}}, TypeState{map[string]Type{"x": TypeBool, "y": TypeInt}}, TypeState{map[string]Type{"x": TypeBool, "y": TypeInt}}, true, false},
}

func TestIfThenElseCheck(t *testing.T) {
	for _, test := range testIfThenElseCheckTests {
		if (test.input.Check(&test.ts) == test.want2 && reflect.DeepEqual(test.ts, test.want1)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(test.ts), StructToJson(test.want1), test.compliant)
		}
	}
}
