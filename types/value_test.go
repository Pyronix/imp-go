package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestIntValue tests the IntValue function

type TestIntValueCase struct {
	input     int
	want      Value
	compliant bool
}

var testIntValueTests = []TestIntValueCase{
	{0, Value{ValueType: ValueInt, IntValue: 0}, true},
	{0, Value{ValueType: ValueInt, IntValue: 1}, false},
}

func TestIntValue(t *testing.T) {
	for _, test := range testIntValueTests {
		if got := IntValue(test.input); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestBoolValue tests the BoolValue function

type TestBoolValueCase struct {
	input     bool
	want      Value
	compliant bool
}

var testBoolValueTests = []TestBoolValueCase{
	{true, Value{ValueType: ValueBool, BoolValue: true}, true},
	{true, Value{ValueType: ValueBool, BoolValue: false}, false},
}

func TestBoolValue(t *testing.T) {
	for _, test := range testBoolValueTests {
		if got := BoolValue(test.input); got != test.want && test.compliant {
			t.Errorf("got %s not equal to want %s", StructToJson(got), StructToJson(test.want))
		}
	}
}

// TestUndefinedValue tests the UndefinedValue function

type TestUndefinedValueCase struct {
	want      Value
	compliant bool
}

var testUndefinedValueTests = []TestUndefinedValueCase{
	{Value{ValueType: Undefined}, true},
	{Value{ValueType: ValueInt}, false},
}

func TestUndefinedValue(t *testing.T) {
	for _, test := range testUndefinedValueTests {
		if got := UndefinedValue(); got != test.want && test.compliant {
			t.Errorf("got %s not equal to want %s", StructToJson(got), StructToJson(test.want))
		}
	}
}

// TestShowVal tests the ShowVal function

type TestShowValCase struct {
	input     Value
	want      string
	compliant bool
}

var testShowValTests = []TestShowValCase{
	{Value{ValueType: ValueInt, IntValue: 0}, "0", true},
	{Value{ValueType: ValueInt, IntValue: 1}, "1", true},
	{Value{ValueType: ValueBool, BoolValue: true}, "true", true},
	{Value{ValueType: ValueBool, BoolValue: false}, "false", true},
	{Value{ValueType: Undefined}, "Undefined", true},

	{Value{ValueType: ValueInt, IntValue: 0}, "1", false},
	{Value{ValueType: ValueInt, IntValue: 1}, "0", false},
	{Value{ValueType: ValueBool, BoolValue: true}, "false", false},
	{Value{ValueType: ValueBool, BoolValue: false}, "true", false},
	{Value{ValueType: Undefined}, "Defined", false},
}

func TestShowVal(t *testing.T) {
	for _, test := range testShowValTests {
		if got := ShowVal(test.input); got != test.want && test.compliant {
			t.Errorf("got %s not equal to want %s", StructToJson(got), StructToJson(test.want))
		}
	}
}

// TestDeclare tests the Declare function
type TestValueDeclareCase struct {
	inputVs     ValueState
	inputString string
	inputValue  Value
	want        ValueState

	compliant bool
}

var testValueDeclareTests = []TestValueDeclareCase{
	{ValueState{map[string]Value{}}, "x", Value{ValueType: ValueInt, IntValue: 0}, ValueState{map[string]Value{"x": {ValueType: ValueInt, IntValue: 0}}}, true},
	{ValueState{map[string]Value{}}, "x", Value{ValueType: ValueInt, IntValue: 1}, ValueState{map[string]Value{"x": {ValueType: ValueInt, IntValue: 0}}}, false},
}

func TestValueDeclare(t *testing.T) {
	for _, test := range testValueDeclareTests {
		test.inputVs.Declare(test.inputString, test.inputValue)
		if reflect.DeepEqual(test.inputVs, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s", StructToJson(test.inputVs), StructToJson(test.want))
		}
	}
}

// TestAssign tests the Assign function
type TestValueAssignCase struct {
	inputVs     ValueState
	inputString string
	inputValue  Value
	want        ValueState
	panic       bool
	compliant   bool
}

var testValueAssignTests = []TestValueAssignCase{
	{ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, "x", Value{ValueType: ValueInt, IntValue: 1}, ValueState{map[string]Value{"x": {ValueType: ValueInt, IntValue: 1}}}, false, true},
	{ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, "x", Value{ValueType: ValueBool, BoolValue: true}, ValueState{map[string]Value{"x": {ValueType: ValueInt, IntValue: 0}}}, true, true},

	{ValueState{map[string]Value{}}, "x", Value{ValueType: ValueBool, BoolValue: true}, ValueState{map[string]Value{}}, true, true},
}

func TestValueAssign(t *testing.T) {
	for _, test := range testValueAssignTests {
		func() {
			defer func() { _ = recover() }()
			test.inputVs.Assign(test.inputString, test.inputValue)
			if reflect.DeepEqual(test.inputVs, test.want) != test.compliant || test.panic != (recover() != nil) {
				t.Errorf("got %s not equal to want %s", StructToJson(test.inputVs), StructToJson(test.want))
			}
		}()
	}
}

// TestPushValueScope tests the PushValueScope function
type TestPushValueScopeCase struct {
	inputVs    ValueState
	want       ValueState
	wantLength int
	compliant  bool
}

var testPushValueScopeTests = []TestPushValueScopeCase{

	{ValueState{}, ValueState{map[string]Value{}, map[string]Value{}}, 2, true},
	{ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, ValueState{map[string]Value{"x": {ValueInt, 0, false}}, map[string]Value{}}, 2, true},
}

func TestPushValueScope(t *testing.T) {
	for _, test := range testPushValueScopeTests {
		PushValueScope(&test.inputVs)
		if (reflect.DeepEqual(test.inputVs, test.want) && len(test.inputVs) == test.wantLength) != test.compliant {
			t.Errorf("got %s not equal to want %s, length from got %d length from want %d", StructToJson(test.inputVs), StructToJson(test.want), len(test.inputVs), test.wantLength)
		}
	}
}

// TestPopValueScope tests the PopValueScope function
type TestPopValueScopeCase struct {
	inputVs    ValueState
	want       ValueState
	wantLength int
	panic      bool
	compliant  bool
}

var testPopValueScopeTests = []TestPopValueScopeCase{

	{ValueState{map[string]Value{}, map[string]Value{}}, ValueState{map[string]Value{}}, 1, false, true},
	{ValueState{map[string]Value{}}, ValueState{map[string]Value{}}, 1, true, true},
	{ValueState{map[string]Value{"x": {ValueInt, 0, false}}, map[string]Value{}}, ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, 1, false, true},
}

func TestPopValueScope(t *testing.T) {
	for _, test := range testPopValueScopeTests {
		func() {
			defer func() { _ = recover() }()
			PopValueScope(&test.inputVs)
			if (reflect.DeepEqual(test.inputVs, test.want) && len(test.inputVs) == test.wantLength) != test.compliant || test.panic != (recover() != nil) {
				t.Errorf("got %s not equal to want %s, length from got %d length from want %d", StructToJson(test.inputVs), StructToJson(test.want), len(test.inputVs), test.wantLength)
			}
		}()
	}
}

// TestGetCurrentValueScope tests the GetCurrentValueScope function
type TestGetCurrentValueScopeCase struct {
	inputVs    ValueState
	want       map[string]Value
	wantLength int
	compliant  bool
}

var testGetCurrentValueScopeTests = []TestGetCurrentValueScopeCase{

	{ValueState{}, map[string]Value{}, 1, true},
	{ValueState{map[string]Value{}}, map[string]Value{}, 1, true},
	{ValueState{map[string]Value{"x": {ValueInt, 0, false}}, map[string]Value{}}, map[string]Value{}, 1, true},
}

func TestGetCurrentValueScope(t *testing.T) {
	for _, test := range testGetCurrentValueScopeTests {
		got := test.inputVs.GetCurrentValueScope()
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s", StructToJson(got), StructToJson(test.want))
		}
	}
}

// LookUpValueByVariableName tests the LookUpValueByVariableName function
type TestLookUpValueByVariableNameCase struct {
	inputTs   ValueState
	typeName  string
	wantValue Value
	wantBool  bool
	compliant bool
}

var testLookUpValueByVariableNameTests = []TestLookUpValueByVariableNameCase{
	{ValueState{map[string]Value{"x": {ValueInt, 1, false}}}, "x", Value{ValueInt, 1, false}, true, true},
	{ValueState{map[string]Value{"x": {ValueInt, 1, false}}, map[string]Value{"x": {ValueBool, 0, true}}}, "x", Value{ValueBool, 0, true}, true, true},
	{ValueState{map[string]Value{"x": {ValueInt, 1, false}}, map[string]Value{}}, "x", Value{ValueInt, 1, false}, true, true},
	{ValueState{map[string]Value{"x": {ValueInt, 0, false}}}, "y", Value{}, false, true},
}

func TestLookUpValueByVariableNameScope(t *testing.T) {
	for _, test := range testLookUpValueByVariableNameTests {
		gotType, gotBool := test.inputTs.LookUpValueByVariableName(test.typeName)
		if (gotType == test.wantValue && gotBool == test.wantBool) != test.compliant {
			t.Errorf("gotType %s not equal to wantType %s or gotBool %s not equal to wantBool %s", StructToJson(gotType), StructToJson(test.wantValue), StructToJson(gotBool), StructToJson(test.wantBool))
		}
	}
}
