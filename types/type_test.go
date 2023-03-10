package types

import (
	. "imp/helper"
	"reflect"
	"testing"
)

// TestShowType tests the ShowType function

type TestShowTypeCase struct {
	input     Type
	want      string
	compliant bool
}

var testShowTypeTests = []TestShowTypeCase{
	{TypeIllTyped, "Illtyped", true},
	{TypeInt, "Int", true},
	{TypeBool, "BoolExpression", true},

	{TypeIllTyped, "Int", false},
	{TypeInt, "BoolExpression", false},
	{TypeBool, "Illtyped", false},
}

func TestShowType(t *testing.T) {
	for _, test := range testShowTypeTests {
		if got := ShowType(test.input); (reflect.DeepEqual(got, test.want)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test should be %t", StructToJson(got), StructToJson(test.want), test.compliant)
		}
	}
}

// TestDeclare tests the Declare function
type TestTypeDeclareCase struct {
	inputTs     TypeState
	inputString string
	inputType   Type
	want        TypeState

	compliant bool
}

var testTypeDeclareTests = []TestTypeDeclareCase{
	{TypeState{map[string]Type{}}, "x", TypeInt, TypeState{map[string]Type{"x": TypeInt}}, true},
	{TypeState{map[string]Type{"x": TypeInt}}, "y", TypeInt, TypeState{map[string]Type{"x": TypeInt, "y": TypeInt}}, true},
	{TypeState{map[string]Type{}}, "x", TypeBool, TypeState{map[string]Type{"x": TypeBool}}, true},
}

func TestTypeDeclare(t *testing.T) {
	for _, test := range testTypeDeclareTests {
		test.inputTs.Declare(test.inputString, test.inputType)
		if reflect.DeepEqual(test.inputTs, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s", StructToJson(test.inputTs), StructToJson(test.want))
		}
	}
}

// TestAssign tests the Assign function
type TestTypeAssignCase struct {
	inputTs     TypeState
	inputString string
	inputType   Type
	want        TypeState
	panic       bool
	compliant   bool
}

var testTypeAssignTests = []TestTypeAssignCase{
	{TypeState{map[string]Type{"x": TypeInt}}, "x", TypeInt, TypeState{map[string]Type{"x": TypeInt}}, false, true},
	{TypeState{map[string]Type{"x": TypeInt}}, "x", TypeBool, TypeState{map[string]Type{"x": TypeInt}}, true, true},
	{TypeState{map[string]Type{}}, "x", TypeBool, TypeState{map[string]Type{}}, true, true},
}

func TestTypeAssign(t *testing.T) {
	for _, test := range testTypeAssignTests {

		if test.panic {
			func() {
				defer func() { _ = recover() }()
				test.inputTs.Assign(test.inputString, test.inputType)
				t.Errorf("expected panic but it is not")
			}()
		} else {
			test.inputTs.Assign(test.inputString, test.inputType)
			if reflect.DeepEqual(test.inputTs, test.want) != test.compliant {
				t.Errorf("got %s not equal to want %s", StructToJson(test.inputTs), StructToJson(test.want))
			}
		}
	}
}

// TestPushTypeScope tests the PushTypeScope function
type TestPushTypeScopeCase struct {
	inputTs    TypeState
	want       TypeState
	wantLength int
	compliant  bool
}

var testPushTypeScopeTests = []TestPushTypeScopeCase{

	{TypeState{}, TypeState{map[string]Type{}, map[string]Type{}}, 2, true},
	{TypeState{map[string]Type{"x": TypeInt}}, TypeState{map[string]Type{"x": TypeInt}, map[string]Type{}}, 2, true},
}

func TestPushTypeScope(t *testing.T) {
	for _, test := range testPushTypeScopeTests {
		PushTypeScope(&test.inputTs)
		if (reflect.DeepEqual(test.inputTs, test.want) && len(test.inputTs) == test.wantLength) != test.compliant {
			t.Errorf("got %s not equal to want %s, length from got %d length from want %d", StructToJson(test.inputTs), StructToJson(test.want), len(test.inputTs), test.wantLength)
		}
	}
}

// TestPopTypeScope tests the PopTypeScope function
type TestPopTypeScopeCase struct {
	inputTs    TypeState
	want       TypeState
	wantLength int
	panic      bool
	compliant  bool
}

var testPopTypeScopeTests = []TestPopTypeScopeCase{

	{TypeState{map[string]Type{}, map[string]Type{}}, TypeState{map[string]Type{}}, 1, false, true},
	{TypeState{map[string]Type{}}, TypeState{map[string]Type{}}, 1, true, true},
	{TypeState{map[string]Type{"x": TypeInt}, map[string]Type{}}, TypeState{map[string]Type{"x": TypeInt}}, 1, false, true},
}

func TestTypeTypeScope(t *testing.T) {
	for _, test := range testPopTypeScopeTests {
		if test.panic {
			func() {
				defer func() { _ = recover() }()
				PopTypeScope(&test.inputTs)
				t.Errorf("expected panic but it is not")
			}()
		} else {
			PopTypeScope(&test.inputTs)
			if (reflect.DeepEqual(test.inputTs, test.want) && len(test.inputTs) == test.wantLength) != test.compliant {
				t.Errorf("got %s not equal to want %s, length from got %d length from want %d", StructToJson(test.inputTs), StructToJson(test.want), len(test.inputTs), test.wantLength)
			}
		}
	}
}

// TestGetCurrentTypeScope tests the GetCurrentTypeScope function
type TestGetCurrentTypeScopeCase struct {
	inputTs    TypeState
	want       map[string]Type
	wantLength int
	compliant  bool
}

var testGetCurrentTypeScopeTests = []TestGetCurrentTypeScopeCase{

	{TypeState{}, map[string]Type{}, 1, true},
	{TypeState{map[string]Type{}}, map[string]Type{}, 1, true},
	{TypeState{map[string]Type{"x": TypeInt}, map[string]Type{"x": TypeBool}}, map[string]Type{"x": TypeBool}, 1, true},
}

func TestGetCurrentTypeScope(t *testing.T) {
	for _, test := range testGetCurrentTypeScopeTests {
		got := test.inputTs.GetCurrentTypeScope()
		if reflect.DeepEqual(got, test.want) != test.compliant {
			t.Errorf("got %s not equal to want %s", StructToJson(got), StructToJson(test.want))
		}
	}
}

// LookUpTypeByVariableName tests the LookUpTypeByVariableName function
type TestLookUpTypeByVariableNameCase struct {
	inputTs   TypeState
	typeName  string
	wantType  Type
	wantBool  bool
	compliant bool
}

var testLookUpTypeByVariableNameTests = []TestLookUpTypeByVariableNameCase{
	{TypeState{map[string]Type{"x": TypeInt}, map[string]Type{"x": TypeBool}}, "x", TypeBool, true, true},
	{TypeState{map[string]Type{"x": TypeInt}, map[string]Type{}}, "x", TypeInt, true, true},
	{TypeState{map[string]Type{"x": TypeInt}}, "y", TypeIllTyped, false, true},
}

func TestLookUpTypeByVariableNameScope(t *testing.T) {
	for _, test := range testLookUpTypeByVariableNameTests {
		gotType, gotBool := test.inputTs.LookUpTypeByVariableName(test.typeName)
		if (gotType == test.wantType && gotBool == test.wantBool) != test.compliant {
			t.Errorf("gotType %s not equal to wantType %s or gotBool %s not equal to wantBool %s", StructToJson(gotType), StructToJson(test.wantType), StructToJson(gotBool), StructToJson(test.wantBool))
		}
	}
}
