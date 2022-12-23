package parser

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const numbers = "0123456789"
const lowerLetters = "abcdefghijklmnopqrstuvwxyz"
const upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letters = lowerLetters + upperLetters
const alphaNumeric = letters + numbers

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
func (l *Lexer) takeExactly(str string) bool {
	taken := 0
	for _, char := range str {
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
		switch next := l.next(); {
		case next == EOFRune:
			l.emit(EOF)
			return nil
		case unicode.IsSpace(next):
			l.ignore()
		case next == '{':
			l.emit(BLOCKOPEN)
		case next == '}':
			l.emit(BLOCKCLOSE)
		case next == ';':
			l.emit(SEMICOLON)
		case next == ':' && l.take("="):
			l.emit(DECLARATION)
		case next == '=' && l.take("="):
			l.emit(EQUAL)
		case next == '=':
			l.emit(ASSIGMENT)
		case next == '+':
			l.emit(ADD)
		case next == '*':
			l.emit(MUL)
		case next == '|' && l.take("|"):
			l.emit(OR)
		case next == '&' && l.take("&"):
			l.emit(AND)
		case next == '!':
			l.emit(NOT)
		case next == '<':
			l.emit(LESS)
		case next == '(':
			l.emit(OPEN)
		case next == ')':
			l.emit(CLOSE)
		case strings.ContainsRune(lowerLetters, next):
			l.rewind()
			return lexWord
		case next == '-' || strings.ContainsRune(numbers, next):
			l.rewind()
			return lexInt
		default:
			l.emit(ILLEGAL)
		}
	}
}

func lexWord(l *Lexer) LexerStateFunc {
	// check for reserved words
	if l.takeExactly("while") && !strings.ContainsRune(alphaNumeric+"_", l.peek()) {
		l.emit(WHILE)
		return lexCode
	}
	if l.takeExactly("if") && !strings.ContainsRune(alphaNumeric+"_", l.peek()) {
		l.emit(IF)
		return lexCode
	}
	if l.takeExactly("else") && !strings.ContainsRune(alphaNumeric+"_", l.peek()) {
		l.emit(ELSE)
		return lexCode
	}
	if l.takeExactly("print") && !strings.ContainsRune(alphaNumeric+"_", l.peek()) {
		l.emit(PRINT)
		return lexCode
	}
	if l.takeExactly("true") && !strings.ContainsRune(alphaNumeric+"_", l.peek()) {
		l.emit(BOOL)
		return lexCode
	}
	if l.takeExactly("false") && !strings.ContainsRune(alphaNumeric+"_", l.peek()) {
		l.emit(BOOL)
		return lexCode
	}

	return lexIdentifier
}

func lexIdentifier(l *Lexer) LexerStateFunc {
	l.takeMany(alphaNumeric + "_")
	l.emit(IDENTIFIER)

	return lexCode
}

func lexInt(l *Lexer) LexerStateFunc {
	l.take("-")
	l.takeMany(numbers)

	if strings.ContainsRune(alphaNumeric, l.peek()) {
		l.next()
		return l.errorf("bad number syntax: %q", l.current())
	}
	l.emit(INT)

	return lexCode
}
