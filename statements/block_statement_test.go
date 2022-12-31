package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestBlock tests the Block function

type TestBlockCase struct {
	input     Statement
	want      BlockStatement
	compliant bool
}

var testBlockTests = []TestBlockCase{
	{DeclarationStatement{"x", NumberExpression(1)}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, true},
	{DeclarationStatement{"x", NumberExpression(1)}, BlockStatement{DeclarationStatement{"x", NumberExpression(2)}}, false},
}

func TestBlock(t *testing.T) {
	for _, test := range testBlockTests {
		got := BlockStatement{test.input}
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestBlockEval tests the Eval function
type TestBlockEvalCase struct {
	input     BlockStatement
	want      ValueState
	compliant bool
}

var testBlockEvalTests = []TestBlockEvalCase{
	{BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 1, false}}, true},
	{BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, ValueState{"x": Value{ValueInt, 2, false}}, false},
	{BlockStatement{DeclarationStatement{"x", BoolExpression(true)}}, ValueState{"x": Value{ValueInt, 0, true}}, false},
}

func TestBlockEval(t *testing.T) {
	for _, test := range testBlockEvalTests {
		got := ValueState{"x": Value{ValueInt, 1, false}}

		test.input.Eval(got)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestBlockPretty tests the Pretty function

type TestBlockPrettyCase struct {
	input     BlockStatement
	want      string
	compliant bool
}

var testBlockPrettyTests = []TestBlockPrettyCase{
	{BlockStatement{DeclarationStatement{"x", BoolExpression(true)}}, "{x := true}", true},
	{BlockStatement{DeclarationStatement{"x", BoolExpression(true)}}, "x := true", false},
}

func TestBlockPretty(t *testing.T) {
	for _, test := range testBlockPrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestBlockCheck tests the Check function

type TestBlockCheckCase struct {
	input1    TypeState
	input2    BlockStatement
	want1     TypeState
	want2     bool
	compliant bool
}

var testBlockCheckTests = []TestBlockCheckCase{
	{TypeState{"x": TypeInt}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, TypeState{"x": TypeInt}, true, true},
	{TypeState{"x": TypeInt}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, TypeState{"x": TypeInt}, false, false},
	{TypeState{"x": TypeInt}, BlockStatement{DeclarationStatement{"x", NumberExpression(1)}}, TypeState{"x": TypeBool}, true, false},
}

func TestBlockCheck(t *testing.T) {
	for _, test := range testBlockCheckTests {
		got := test.input1
		if (reflect.DeepEqual(got, test.want1) && test.input2.Check(got) == test.want2) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want1), test.compliant)
		}
	}
}
