package parser

import (
	. "imp/helper"
	. "imp/statements"
	. "imp/types"
	"io"
	"reflect"
	"strings"
	"testing"
)

type TestParserCase struct {
	inputTokens   *Tape[Token]
	inputPosition int
	wantParser    Parser
	compliant     bool
}

var testParserTests = []TestParserCase{
	{TokenizeString("x := 1"), 4, Parser{TokenizeString("x := 1"), 4}, true},
	{TokenizeString("while"), 4, Parser{TokenizeString("while"), 4}, true},
	{TokenizeString("while"), 0, Parser{TokenizeString("if"), 0}, false},
}

func TestParser(t *testing.T) {
	for _, test := range testParserTests {
		got := Parser{test.inputTokens, test.inputPosition}
		if (reflect.DeepEqual(got, test.wantParser)) != test.compliant {
			t.Errorf("got %s not equal to want %s, test was %t but want %t", StructToJson(got), StructToJson(test.wantParser), reflect.DeepEqual(got, test.wantParser), test.compliant)
		}
	}
}

type TestNewParserCase struct {
	input      string
	wantParser Parser
	compliant  bool
}

var testNewParserTests = []TestNewParserCase{
	{"x := 1", Parser{TokenizeString("x := 1"), 0}, true},
	{"while", Parser{TokenizeString("while"), 0}, true},
	{"x := 1", Parser{TokenizeString("x := 2"), 0}, false},
	{"x := 1", Parser{TokenizeString("x := 1"), 2}, false},
}

func TestNewParser(t *testing.T) {
	for _, test := range testNewParserTests {
		got := *NewParser(test.input)
		if (reflect.DeepEqual(got.tokens, test.wantParser.tokens) && got.position == test.wantParser.position) != test.compliant {
			t.Errorf("got %s not equal to want %s, test was %t but want %t", StructToJson(got), StructToJson(test.wantParser), reflect.DeepEqual(got, test.wantParser), test.compliant)
		}
	}
}

type TestNewParseFromReaderCase struct {
	input      io.Reader
	wantParser Parser
	compliant  bool
}

var testNewParserFromReaderTests = []TestNewParseFromReaderCase{
	{strings.NewReader("x := 1"), Parser{TokenizeString("x := 1"), 0}, true},
	{strings.NewReader("while"), Parser{TokenizeString("while"), 0}, true},
	{strings.NewReader("x := 1"), Parser{TokenizeString("x := 2"), 0}, false},
	{strings.NewReader("x := 1"), Parser{TokenizeString("x := 1"), 2}, false},
}

func TestNewParserFromReader(t *testing.T) {
	for _, test := range testNewParserFromReaderTests {
		got := *NewParserFromReader(test.input)
		if (reflect.DeepEqual(got.tokens, test.wantParser.tokens) && got.position == test.wantParser.position) != test.compliant {
			t.Errorf("got %s not equal to want %s, test was %t but want %t", StructToJson(got), StructToJson(test.wantParser), reflect.DeepEqual(got, test.wantParser), test.compliant)
		}
	}
}

type TestParseProgramCase struct {
	input       *Parser
	wantProgram ProgramStatement
	panic       bool
	compliant   bool
}

var testParseProgramTests = []TestParseProgramCase{
	{NewParser("x := 1"), Program(Declaration("x", NumberExpression(1))), false, true},
	{NewParser("x := 0; y := 1; while x < 10 {x = x + 1; y := true}"), Program(SequenceStatement{Declaration("x", Number(0)), Sequence(Declaration("y", Number(1)), While(Lesser(Variable("x"), Number(10)), Block(Sequence(Assignment("x", Plus(Variable("x"), Number(1))), Declaration("y", Bool(true))))))}), false, true},
	{NewParser("x"), Program(Declaration("x", Number(1))), true, false},
	{NewParser("x = 1"), Program(Assignment("x", NumberExpression(1))), false, true},
	{NewParser("if 1 < 2 {x := 1} else {y := 2}"), Program(Ite(Lesser(Number(1), Number(2)), Block(Declaration("x", Number(1))), Block(Declaration("y", Number(2))))), false, true},
	{NewParser("print 1"), Program(Print(Number(1))), false, true},
	{NewParser("if true || false {x := false} else {y := 2}"), Program(Ite(Or(Bool(true), Bool(false)), Block(Declaration("x", Bool(false))), Block(Declaration("y", Number(2))))), false, true},
	{NewParser("if true && false {x := 1} else {y := 2}"), Program(Ite(And(Bool(true), Bool(false)), Block(Declaration("x", Number(1))), Block(Declaration("y", Number(2))))), false, true},
	{NewParser("if 1 == 2 {x := 1} else {y := 2}"), Program(Ite(Equal(Number(1), Number(2)), Block(Declaration("x", Number(1))), Block(Declaration("y", Number(2))))), false, true},
	{NewParser("x := 1 * 2"), Program(Declaration("x", Mult(Number(1), Number(2)))), false, true},

	{NewParser("x := !true"), Program(Declaration("x", Negation(Bool(true)))), false, true},
	{NewParser("x := (1 + 2) * 2"), Program(Declaration("x", Mult(Grouping(Plus(Number(1), Number(2))), Number(2)))), false, true},

	{NewParser("if 1 < 2 {x := 1} { y := 2}"), Program(Ite(Lesser(Number(1), Number(2)), Block(Declaration("x", Number(1))), Block(Declaration("y", Number(2))))), true, false},
	{NewParser("if 1 < 2 x := 1"), Program(Ite(Lesser(Number(1), Number(2)), Block(Declaration("x", Number(1))), Block(Declaration("y", Number(2))))), true, false},
	{NewParser("if 1 < 2 {x := 1} else {y := 2"), Program(Ite(Lesser(Number(1), Number(2)), Block(Declaration("x", Number(1))), Block(Declaration("y", Number(2))))), true, false},
	{NewParser("1"), Program(Print(Number(1))), true, false},
	{NewParser("x := 2.5"), Program(Declaration("x", Number(1))), true, false},
	{NewParser("x := !(1"), Program(Declaration("x", Number(1))), true, false},
	{NewParser("x := 18446744073709551616"), Program(Declaration("x", Number(1))), true, false},
	{NewParser("x := if 2 < 1 {x := 1} else{ y := 2}"), Program(Declaration("x", Number(1))), true, false},
	{NewParser("x"), Program(Declaration("x", Number(1))), true, false},
	{&Parser{TokenizeString("x := 1"), -1}, Program(Declaration("x", Number(1))), false, true},
}

func TestParseProgram(t *testing.T) {
	for _, test := range testParseProgramTests {
		if test.panic {
			func() {
				defer func() { recover() }()
				(*test.input).ParseProgram()
				t.Errorf("expected ParseProgram() to panic")
			}()
		} else {
			got := (*test.input).ParseProgram()
			if (reflect.DeepEqual(got, test.wantProgram)) != test.compliant {
				t.Errorf("got %s not equal to want %s, test was %t but want %t", StructToJson(got), StructToJson(test.wantProgram), reflect.DeepEqual(got, test.wantProgram), test.compliant)
			}
		}
	}
}

// direct test to cover edge cases
func TestParsePrimary(t *testing.T) {
	var parser Parser

	parser = Parser{
		tokens: &Tape[Token]{
			data: []Token{
				{
					0,
					BOOL,
					"not_true_or_false",
				},
			},
			eofValue: Token{},
			size:     1,
			position: 0,
		},
		position: 0,
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("expected parsePrimary() to panic when receiving an invalid bool literal")
		}
	}()

	parser.parsePrimary()
}
