package statements

import (
	. "imp/helper"
	. "imp/types"
	"reflect"
	"testing"
)

// TestPrint tests the Print function

type TestPrintCase struct {
	input     Expression
	want      PrintStatement
	compliant bool
}

var testPrintTests = []TestPrintCase{
	{NumberExpression(1), PrintStatement{NumberExpression(1)}, true},
	{NumberExpression(1), PrintStatement{NumberExpression(2)}, false},
	{PlusExpression{NumberExpression(1), NumberExpression(2)}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, true},
	{PlusExpression{NumberExpression(1), NumberExpression(2)}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(3)}}, false},
}

func TestPrint(t *testing.T) {
	for _, test := range testPrintTests {
		got := Print(test.input)
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestPrintEval tests the Eval function
type TestPrintEvalCase struct {
	input     PrintStatement
	vs        ValueState
	want      ValueState
	compliant bool
}

var testPrintEvalTests = []TestPrintEvalCase{
	{PrintStatement{NumberExpression(1)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, true},
	{PrintStatement{NumberExpression(1)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 2, false}}}, false},
	{PrintStatement{NumberExpression(1)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 1, true}}}, false},
	{PrintStatement{NumberExpression(1)}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}, "y": {ValueInt, 1, false}}}, false},
	{PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, true},
	{PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 2, false}}}, false},
	{PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 1, true}}}, false},
	{PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, ValueState{map[string]Value{"x": {ValueInt, 1, false}, "y": {ValueInt, 1, false}}}, false},
}

func TestPrintEval(t *testing.T) {
	for _, test := range testPrintEvalTests {
		test.input.Eval(&test.vs)
		if reflect.DeepEqual(test.vs, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(test.vs), StructToJson(test.want), test.compliant)
		}
	}
}

// TestPrintPretty tests the Pretty function

type TestPrintPrettyCase struct {
	input     PrintStatement
	want      string
	compliant bool
}

var testPrintPrettyTests = []TestPrintPrettyCase{
	{PrintStatement{NumberExpression(1)}, "print 1", true},
	{PrintStatement{NumberExpression(1)}, "print 2", false},
	{PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, "print 1 + 2", true},
	{PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, "print 1 + 3", false},
}

func TestPrintPretty(t *testing.T) {
	for _, test := range testPrintPrettyTests {
		if got := test.input.Pretty(); reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", got, test.want, test.compliant)
		}
	}
}

// TestPrintCheck tests the Check function

type TestPrintCheckCase struct {
	input1    TypeState
	input2    PrintStatement
	want1     TypeState
	want2     bool
	compliant bool
}

var testPrintCheckTests = []TestPrintCheckCase{
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{NumberExpression(1)}, TypeState{map[string]Value{"x": TypeInt}}, true, true},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{NumberExpression(1)}, TypeState{map[string]Value{"x": TypeInt}}, false, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{NumberExpression(1)}, TypeState{map[string]Value{"x": TypeBool}}, true, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{NumberExpression(1)}, TypeState{map[string]Value{"x": TypeBool}}, false, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{NumberExpression(1)}, TypeState{map[string]Value{"x": TypeInt, "y": TypeInt}}, true, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{NumberExpression(1)}, TypeState{map[string]Value{"x": TypeInt, "y": TypeInt}}, false, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, TypeState{map[string]Value{"x": TypeInt}}, true, true},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, TypeState{map[string]Value{"x": TypeInt}}, false, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, TypeState{map[string]Value{"x": TypeBool}}, true, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, TypeState{map[string]Value{"x": TypeBool}}, false, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, TypeState{map[string]Value{"x": TypeInt, "y": TypeInt}}, true, false},
	{TypeState{map[string]Value{"x": TypeInt}}, PrintStatement{PlusExpression{NumberExpression(1), NumberExpression(2)}}, TypeState{map[string]Value{"x": TypeInt, "y": TypeInt}}, false, false},
}

func TestPrintCheck(t *testing.T) {
	for _, test := range testPrintCheckTests {
		got := test.input1
		if (reflect.DeepEqual(got, test.want1) && test.input2.Check(&got) == test.want2) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want1), test.compliant)
		}
	}
}
