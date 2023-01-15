package parser

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

const numbers = "0123456789"
const lowerLetters = "abcdefghijklmnopqrstuvwxyz"
const upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letters = lowerLetters + upperLetters
const alphaNumeric = letters + numbers

type Tokenizer struct {
	runes      *Tape[rune]
	tokenStart int
	tokens     chan Token
}

type TokenizerStateFunc func(*Tokenizer) TokenizerStateFunc

func TokenizeString(input string) *Tape[Token] {
	return TokenizeFromReader(strings.NewReader(input))
}

func TokenizeFromReader(input io.Reader) *Tape[Token] {
	tokenizer := Tokenizer{
		runes:      NewTapeFromReader(input),
		tokenStart: 0,
		tokens:     make(chan Token),
	}
	go tokenizer.run()

	var tokens []Token
	lastToken := <-tokenizer.tokens

	for lastToken.Type != EOF {
		tokens = append(tokens, lastToken)
		lastToken = <-tokenizer.tokens
	}

	return &Tape[Token]{
		data:     tokens,
		eofValue: lastToken,
		size:     len(tokens),
	}
}

func (l *Tokenizer) run() {
	for state := tokenizeCode; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func (l *Tokenizer) currentValue() string {
	var runes []rune
	runes = l.runes.ReadSlice(l.tokenStart, l.runes.position)

	var nonEofRunes []rune
	for i := 0; i < len(runes); i++ {
		if runes[i] != EOFRune {
			nonEofRunes = append(nonEofRunes, runes[i])
		}
	}

	return string(nonEofRunes)
}

func (l *Tokenizer) emit(t TokenType) {
	l.tokens <- Token{
		l.tokenStart,
		t,
		l.currentValue(),
	}
	l.tokenStart = l.runes.position
}

func (l *Tokenizer) ignore() {
	l.tokenStart = l.runes.position
}

// takes one valid rune
func (l *Tokenizer) take(valid string) bool {
	if strings.ContainsRune(valid, l.runes.Next()) {
		return true
	}
	l.runes.Rewind()

	return false
}

// takes multiple valid runes, at least one
func (l *Tokenizer) takeMany(valid string) bool {
	if !l.take(valid) {
		return false
	}

	for l.take(valid) {
	}

	return true
}

// takes a full literal exactly
func (l *Tokenizer) takeExactly(str string) bool {
	taken := 0
	for _, char := range str {
		if l.take(string(char)) {
			taken++
		} else {
			for i := 0; i < taken; i++ {
				l.runes.Rewind()
			}

			return false
		}
	}

	return true
}

func (l *Tokenizer) errorf(format string, args ...interface{}) TokenizerStateFunc {
	return func(l *Tokenizer) TokenizerStateFunc {
		l.tokens <- Token{
			l.tokenStart,
			ERROR,
			fmt.Sprintf(format, args...),
		}

		return nil
	}
}

func tokenizeCode(l *Tokenizer) TokenizerStateFunc {
	for {
		switch next := l.runes.Next(); {
		case next == EOFRune:
			l.emit(EOF)
			return nil
		case unicode.IsSpace(next):
			l.ignore()
		case next == '/' && l.runes.Peek() == '/':
			// ignore comments, they start with '//' and end with a new line
			for l.runes.Peek() != '\n' {
				l.runes.Next()
			}
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
			l.runes.Rewind()
			return tokenizeWord
		case next == '-' || strings.ContainsRune(numbers, next):
			l.runes.Rewind()
			return tokenizeInt
		default:
			l.emit(ILLEGAL)
		}
	}
}

func tokenizeWord(l *Tokenizer) TokenizerStateFunc {
	// check for reserved words
	if l.takeExactly("while") && !strings.ContainsRune(alphaNumeric+"_", l.runes.Peek()) {
		l.emit(WHILE)
		return tokenizeCode
	}
	if l.takeExactly("if") && !strings.ContainsRune(alphaNumeric+"_", l.runes.Peek()) {
		l.emit(IF)
		return tokenizeCode
	}
	if l.takeExactly("else") && !strings.ContainsRune(alphaNumeric+"_", l.runes.Peek()) {
		l.emit(ELSE)
		return tokenizeCode
	}
	if l.takeExactly("print") && !strings.ContainsRune(alphaNumeric+"_", l.runes.Peek()) {
		l.emit(PRINT)
		return tokenizeCode
	}
	if l.takeExactly("true") && !strings.ContainsRune(alphaNumeric+"_", l.runes.Peek()) {
		l.emit(BOOL)
		return tokenizeCode
	}
	if l.takeExactly("false") && !strings.ContainsRune(alphaNumeric+"_", l.runes.Peek()) {
		l.emit(BOOL)
		return tokenizeCode
	}

	return tokenizeIdentifier
}

func tokenizeIdentifier(l *Tokenizer) TokenizerStateFunc {
	l.takeMany(alphaNumeric + "_")
	l.emit(IDENTIFIER)

	return tokenizeCode
}

func tokenizeInt(l *Tokenizer) TokenizerStateFunc {
	l.take("-")
	l.takeMany(numbers)

	if strings.ContainsRune(alphaNumeric, l.runes.Peek()) {
		l.runes.Next()
		return l.errorf("bad number syntax: %q", l.currentValue())
	}
	l.emit(INT)

	return tokenizeCode
}
