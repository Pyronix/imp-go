package parser

import (
	"strings"
	"testing"
)

type nextTest struct {
	input          string
	expectedTokens []Token
}

var nextTests = []nextTest{
	// Tests for EOF
	{" ", []Token{
		{1, EOF, ""},
	}},

	// Tests for ILLEGAL
	{" '", []Token{
		{1, ILLEGAL, "'"},
		{2, EOF, ""},
	}},
	{"\"", []Token{
		{0, ILLEGAL, "\""},
		{1, EOF, ""},
	}},

	// Test for IDENTIFIER
	{"myVar", []Token{
		{0, IDENTIFIER, "myVar"},
		{5, EOF, ""},
	}},
	{"while_var_13", []Token{
		{0, IDENTIFIER, "while_var_13"},
		{12, EOF, ""},
	}},

	// Test for BLOCKOPEN
	{"{", []Token{
		{0, BLOCKOPEN, "{"},
		{1, EOF, ""},
	}},

	// Test for BLOCKCLOSE
	{"}", []Token{
		{0, BLOCKCLOSE, "}"},
		{1, EOF, ""},
	}},

	// Tests for SEMICOLON
	{";", []Token{
		{0, SEMICOLON, ";"},
		{1, EOF, ""},
	}},

	// Tests for DECLARATION
	{":=", []Token{
		{0, DECLARATION, ":="},
		{2, EOF, ""},
	}},
	{":", []Token{
		{0, ILLEGAL, ":"},
		{1, EOF, ""},
	}},

	// Tests for ASSIGMENT
	{"=", []Token{
		{0, ASSIGMENT, "="},
		{1, EOF, ""},
	}},

	// Tests for WHILE
	{"while", []Token{
		{0, WHILE, "while"},
		{5, EOF, ""},
	}},

	// Tests for IF
	{"if", []Token{
		{0, IF, "if"},
		{2, EOF, ""},
	}},

	// Tests for ELSE
	{"else", []Token{
		{0, ELSE, "else"},
		{4, EOF, ""},
	}},

	// Tests for PRINT
	{"print", []Token{
		{0, PRINT, "print"},
		{5, EOF, ""},
	}},

	// Tests for INT
	{"1", []Token{
		{0, INT, "1"},
		{1, EOF, ""},
	}},
	{"-1", []Token{
		{0, INT, "-1"},
		{2, EOF, ""},
	}},
	{"1234567890", []Token{
		{0, INT, "1234567890"},
		{10, EOF, ""},
	}},
	{"-1234567890", []Token{
		{0, INT, "-1234567890"},
		{11, EOF, ""},
	}},
	{"1a", []Token{
		{0, ERROR, "bad number syntax: \"1a\""},
	}},

	// Tests for BOOL
	{"true", []Token{
		{0, BOOL, "true"},
		{4, EOF, ""},
	}},
	{"false", []Token{
		{0, BOOL, "false"},
		{5, EOF, ""},
	}},
	// Tests for ADD
	{"+", []Token{
		{0, ADD, "+"},
		{1, EOF, ""},
	}},

	// Tests for MUL
	{"*", []Token{
		{0, MUL, "*"},
		{1, EOF, ""},
	}},

	// Tests for OR
	{"||", []Token{
		{0, OR, "||"},
		{2, EOF, ""},
	}},

	// Tests for AND
	{"&&", []Token{
		{0, AND, "&&"},
		{2, EOF, ""},
	}},

	// Tests for NOT
	{"!", []Token{
		{0, NOT, "!"},
		{1, EOF, ""},
	}},

	// Tests for EQUAL
	{"==", []Token{
		{0, EQUAL, "=="},
		{2, EOF, ""},
	}},

	// Tests for LESS
	{"<", []Token{
		{0, LESS, "<"},
		{1, EOF, ""},
	}},

	// Tests for OPEN
	{"(", []Token{
		{0, OPEN, "("},
		{1, EOF, ""},
	}},

	// Tests for CLOSE
	{")", []Token{
		{0, CLOSE, ")"},
		{1, EOF, ""},
	}},

	// Test for multiple tokens
	{"}}{}}", []Token{
		{0, BLOCKCLOSE, "}"},
		{1, BLOCKCLOSE, "}"},
		{2, BLOCKOPEN, "{"},
		{3, BLOCKCLOSE, "}"},
		{4, BLOCKCLOSE, "}"},
		{5, EOF, ""},
	}},
	{"1 + 1", []Token{
		{0, INT, "1"},
		{2, ADD, "+"},
		{4, INT, "1"},
		{5, EOF, ""},
	}},
	{"a:=1+-1;b:=a<3&&true", []Token{
		{0, IDENTIFIER, "a"},
		{1, DECLARATION, ":="},
		{3, INT, "1"},
		{4, ADD, "+"},
		{5, INT, "-1"},
		{7, SEMICOLON, ";"},
		{8, IDENTIFIER, "b"},
		{9, DECLARATION, ":="},
		{11, IDENTIFIER, "a"},
		{12, LESS, "<"},
		{13, INT, "3"},
		{14, AND, "&&"},
		{16, BOOL, "true"},
		{20, EOF, ""},
	}},

	// Tests for Whitespace Characters
	{" {", []Token{
		{1, BLOCKOPEN, "{"},
		{2, EOF, ""},
	}},
	{" \n{", []Token{
		{2, BLOCKOPEN, "{"},
		{3, EOF, ""},
	}},
	{" \n\t{", []Token{
		{3, BLOCKOPEN, "{"},
		{4, EOF, ""},
	}},
	{" {\r\n\t  }", []Token{
		{1, BLOCKOPEN, "{"},
		{7, BLOCKCLOSE, "}"},
		{8, EOF, ""},
	}},
	{" \t\r\t  }", []Token{
		{6, BLOCKCLOSE, "}"},
		{7, EOF, ""},
	}},

	// Tests for comments
	{"1; // this is a comment \n2", []Token{
		{0, INT, "1"},
		{1, SEMICOLON, ";"},
		{25, INT, "2"},
	}},
}

func TestTokenizer(t *testing.T) {
	for _, test := range nextTests {
		tape := TokenizeString(test.input)

		for callNo, expected := range test.expectedTokens {
			actual := tape.Next()

			if actual != expected {
				t.Errorf("Testing %q on call %d: expected %s got %s", test.input, callNo, expected, actual)
			}
		}
	}
}

func TestTakeMany(t *testing.T) {
	tokenizer := Tokenizer{
		runes:      NewTapeFromReader(strings.NewReader("abc")),
		tokenStart: 0,
		tokens:     make(chan Token),
	}

	if tokenizer.takeMany("123") != false {
		t.Errorf("expected takeMany() to return false when no rune could be taken")
	}
}

func TestTakeExactly(t *testing.T) {
	tokenizer := Tokenizer{
		runes:      NewTapeFromReader(strings.NewReader("abc")),
		tokenStart: 0,
		tokens:     make(chan Token),
	}

	if tokenizer.takeExactly("abfe") != false {
		t.Errorf("expected takeExactly() to return false when runes could not be taken")
	}

	if tokenizer.runes.position != 0 {
		t.Errorf("expected takeExactly() to not change rune tape position when failed")
	}
}
