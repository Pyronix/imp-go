package parser

import (
	"bufio"
	"io"
	"strings"
)

type Tape[T any] struct {
	data     []T
	eofValue T
	size     int
	position int
}

const (
	EOFRune = -1
)

func NewTapeFromReader(reader io.Reader) *Tape[rune] {
	var runes []rune

	r := bufio.NewReader(reader)

	var next rune
	var err error

	for next, _, err = r.ReadRune(); err == nil; next, _, err = r.ReadRune() {
		runes = append(runes, next)
	}

	if err != io.EOF {
		panic(err)
	}

	return &Tape[rune]{
		data:     runes,
		eofValue: EOFRune,
		size:     len(runes),
		position: 0,
	}
}

func NewTapeFromString(input string) *Tape[rune] {
	return NewTapeFromReader(strings.NewReader(input))
}

func (t *Tape[T]) ReadSlice(start int, end int) []T {
	if start >= end {
		return []T{}
	}

	var slice []T
	
	for current := start; current < end; current++ {
		if current < 0 || t.size <= current {
			slice = append(slice, t.eofValue)
		} else {
			slice = append(slice, t.data[current])
		}
	}

	return slice
}

func (t *Tape[T]) Position() int {
	return t.position
}

func (t *Tape[T]) Advance() int {
	t.position++

	return t.position
}

func (t *Tape[T]) Backup() int {
	t.position--

	return t.position
}

func (t *Tape[T]) Peek() T {
	if t.position < 0 || t.size <= t.position {
		return t.eofValue
	}

	return t.data[t.position]
}

func (t *Tape[T]) Next() T {
	r := t.Peek()
	t.Advance()

	return r
}

func (t *Tape[T]) Rewind() T {
	t.Backup()

	return t.Peek()
}
