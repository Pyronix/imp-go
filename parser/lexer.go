package parser

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	EOFRune rune      = -1
	EOF     TokenType = iota
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

type Lexer struct {
	input       string
	start       int
	position    int
	rewindStack []rune
	tokens      chan Token
}

type LexerStateFunc func(*Lexer) LexerStateFunc

func Lex(input string) (*Lexer, chan Token) {
	l := &Lexer{
		input:    input,
		start:    0,
		position: 0,
		tokens:   make(chan Token),
	}
	go l.run()
	return l, l.tokens
}

func (l *Lexer) NextToken() Token {
	return <-l.tokens
}

func (l *Lexer) run() {
	for state := lexCode; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func (l *Lexer) current() string {
	return l.input[l.start:l.position]
}

func (l *Lexer) emit(t TokenType) {
	l.tokens <- Token{
		l.start,
		t,
		l.current(),
	}
	l.start = l.position
	l.rewindStack = []rune{}
}

func (l *Lexer) ignore() {
	l.start = l.position
	l.rewindStack = []rune{}
}

func (l *Lexer) peek() rune {
	r := l.next()
	l.rewind()

	return r
}

func (l *Lexer) next() rune {
	if l.position >= len(l.input) {
		l.rewindStack = append(l.rewindStack, EOFRune)
		return EOFRune
	}

	var r rune
	var s int

	r, s = utf8.DecodeRuneInString(l.input[l.position:])

	l.position += s
	l.rewindStack = append(l.rewindStack, r)

	return r
}

func (l *Lexer) rewind() {
	if len(l.rewindStack) == 0 {
		return
	}

	r := l.rewindStack[len(l.rewindStack)-1]
	l.rewindStack = l.rewindStack[:len(l.rewindStack)-1]

	if r == EOFRune {
		return
	}

	size := utf8.RuneLen(r)
	l.position -= size

	if l.position < l.start {
		l.position = l.start
	}
}

// takes one valid rune
func (l *Lexer) take(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.rewind()

	return false
}

// takes multiple valid runes, at least one
func (l *Lexer) takeMany(valid string) bool {
	if !l.take(valid) {
		return false
	}

	for l.take(valid) {
	}

	return true
}

// takes a full literal exactly
func (l *Lexer) takeLiteral(literal string) bool {
	taken := 0
	for _, char := range literal {
		if l.take(string(char)) {
			taken++
		} else {
			for i := 0; i < taken; i++ {
				l.rewind()
			}

			return false
		}
	}

	return true
}

func (l *Lexer) errorf(format string, args ...interface{}) LexerStateFunc {
	return func(l *Lexer) LexerStateFunc {
		l.tokens <- Token{
			l.start,
			ERROR,
			fmt.Sprintf(format, args...),
		}

		return nil
	}
}

func lexCode(l *Lexer) LexerStateFunc {
	for {
		next := l.next()

		if next == EOFRune {
			l.emit(EOF)
			break
		} else if unicode.IsSpace(next) {
			l.ignore()
		} else if next == '{' {
			l.emit(BLOCKOPEN)
		} else if next == '}' {
			l.emit(BLOCKCLOSE)
		} else if next == ';' {
			l.emit(SEMICOLON)
		} else if next == ':' && l.take("=") {
			l.emit(DECLARATION)
		} else if next == '=' && l.take("=") {
			l.emit(EQUAL)
		} else if next == '=' {
			l.emit(ASSIGMENT)
		} else if next == 'w' && l.takeLiteral("hile") {
			l.emit(WHILE)
		} else if next == 'i' && l.takeLiteral("f") {
			l.emit(IF)
		} else if next == 'e' && l.takeLiteral("lse") {
			l.emit(ELSE)
		} else if next == 'p' && l.takeLiteral("rint") {
			l.emit(PRINT)
		} else if next == 't' && l.takeLiteral("rue") {
			l.emit(BOOL)
		} else if next == 'f' && l.takeLiteral("alse") {
			l.emit(BOOL)
		} else if next == '+' {
			l.emit(ADD)
		} else if next == '*' {
			l.emit(MUL)
		} else if next == '|' && l.take("|") {
			l.emit(OR)
		} else if next == '&' && l.take("&") {
			l.emit(AND)
		} else if next == '!' {
			l.emit(NOT)
		} else if next == '<' {
			l.emit(LESS)
		} else if next == '(' {
			l.emit(OPEN)
		} else if next == ')' {
			l.emit(CLOSE)
		} else if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", next) {
			l.rewind()
			return lexIdentifier
		} else if next == '-' || strings.ContainsRune("0123456789", next) {
			l.rewind()
			return lexInt
		} else {
			l.emit(ILLEGAL)
		}
	}

	return nil
}

func lexIdentifier(l *Lexer) LexerStateFunc {
	next := l.next()

	if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", next) {
		l.emit(ILLEGAL)
	}

	l.takeMany("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	l.emit(IDENTIFIER)

	return lexCode
}

func lexInt(l *Lexer) LexerStateFunc {
	l.take("-")
	l.takeMany("0123456789")
	l.emit(INT)

	return lexCode
}
