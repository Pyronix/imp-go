package parser

import "fmt"

type TokenType int

const (
	EOF TokenType = iota
	ERROR
	ILLEGAL

	IDENTIFIER
	BLOCKOPEN
	BLOCKCLOSE
	SEMICOLON
	DECLARATION
	ASSIGMENT
	WHILE
	IF
	ELSE
	PRINT

	INT
	BOOL
	ADD
	MUL
	OR
	AND
	NOT
	EQUAL
	LESS
	OPEN
	CLOSE
)

func (t TokenType) String() string {
	return []string{
		EOF:         "EOF",
		ERROR:       "ERROR",
		ILLEGAL:     "ILLEGAL",
		IDENTIFIER:  "IDENTIFIER",
		BLOCKOPEN:   "{",
		BLOCKCLOSE:  "}",
		SEMICOLON:   ";",
		DECLARATION: ":=",
		ASSIGMENT:   "=",
		WHILE:       "while",
		IF:          "if",
		ELSE:        "else",
		PRINT:       "print",

		INT:   "INT",
		BOOL:  "BOOL",
		ADD:   "+",
		MUL:   "*",
		OR:    "||",
		AND:   "&&",
		NOT:   "!",
		EQUAL: "==",
		LESS:  "<",
		OPEN:  "(",
		CLOSE: ")",
	}[t]
}

type Token struct {
	Position int
	Type     TokenType
	Value    string
}

func (t Token) String() string {
	return fmt.Sprintf("{ Position: %+v, Type: %q, Value: %q}", t.Position, t.Type, t.Value)
}
