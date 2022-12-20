package parser

import (
	"strings"
	"testing"
)

// Todo add tests for repeated calls / multi token stream
type nextTest struct {
	input            string
	expectedPosition Position
	expectedToken    Token
	expectedLiteral  string
}

var nextTests = []nextTest{
	{" ", Position{1, 1}, EOF, ""},
	{" '", Position{1, 2}, ILLEGAL, "'"},
	{"\"", Position{1, 1}, ILLEGAL, "\""},
	{" {", Position{1, 2}, BLOCKOPEN, "{"},
	{" \n{", Position{2, 1}, BLOCKOPEN, "{"},
	{" \n\t{", Position{2, 2}, BLOCKOPEN, "{"},
	{" \r\n\t  }", Position{2, 4}, BLOCKCLOSE, "}"},
	{" \t\r\t  }", Position{1, 7}, BLOCKCLOSE, "}"},
}

func TestNext(t *testing.T) {
	for _, test := range nextTests {
		lexer := NewLexer(strings.NewReader(test.input))
		position, token, literal := lexer.Next()

		if position != test.expectedPosition {
			t.Errorf("Testing \"%q\": Output Position %+v not equal to expected %+v", test.input, position, test.expectedPosition)
		}

		if token != test.expectedToken {
			t.Errorf("Testing \"%q\": Output Token %q not equal to expected %q", test.input, token, test.expectedToken)
		}

		if literal != test.expectedLiteral {
			t.Errorf("Testing \"%q\": Output Literal %q not equal to expected %q", test.input, literal, test.expectedLiteral)
		}
	}
}
