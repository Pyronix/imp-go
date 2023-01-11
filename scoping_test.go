package main

import (
	. "imp/statements"
	. "imp/types"
	"reflect"
	"testing"
)

type WhileScopeCase struct {
	input   Statement
	wantVs  ValueState
	wantTs  TypeState
	wantStr string
}

var testWhileScopeTests = []WhileScopeCase{
	{Program(SequenceStatement{Sequence(Declaration("x", Number(0)), Declaration("y", Number(1))), While(Lesser(Variable("x"), Number(10)), Block(Sequence(Assignment("x", Plus(Variable("x"), Number(1))), Declaration("y", Bool(true)))))}), ValueState{map[string]Value{"x": IntValue(10), "y": IntValue(1)}}, TypeState{map[string]Type{"x": TypeInt, "y": TypeInt}}, "x := 0; y := 1; while x < 10 {x = x + 1; y := true}"},
}

func runTest(e Statement) (ValueState, TypeState, string) {
	vs := []map[string]Value{{}}
	ts := []map[string]Type{{}}
	e.Eval((*ValueState)(&vs))
	e.Check((*TypeState)(&ts))
	pretty := e.Pretty()
	return vs, ts, pretty
}

func TestWhileScoping1(t *testing.T) {
	for _, test := range testWhileScopeTests {
		vs, ts, pretty := runTest(test.input)
		if !reflect.DeepEqual(vs, test.wantVs) {
			t.Errorf("TestWhileScoping1 ValueState: got %v, want %v", vs, test.wantVs)
		}
		if !reflect.DeepEqual(ts, test.wantTs) {
			t.Errorf("TestWhileScoping1 TypeState: got %v, want %v", ts, test.wantTs)
		}
		if pretty != test.wantStr {
			t.Errorf("TestWhileScoping1 Pretty: got %v, want %v", pretty, test.wantStr)
		}
	}
}
