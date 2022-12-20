package parser

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

const (
	EOF = iota
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

func (t Token) String() string {
	return []string{
		EOF:         "EOF",
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

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Next() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}

		l.pos.column++

		switch r {
		case '\n':
			l.pos.line++
			l.pos.column = 0
		case '{':
			return l.pos, BLOCKOPEN, "{"
		case '}':
			return l.pos, BLOCKCLOSE, "}"
		default:
			if unicode.IsSpace(r) {
				continue
			} else {
				return l.pos, ILLEGAL, string(r)
			}
		}
	}
}
