package parser

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

type MockReader struct {
	nextSlice     []byte
	nextByteCount int
	nextError     error
}

func (m *MockReader) Read(b []byte) (n int, err error) {
	copy(b, m.nextSlice)
	return m.nextByteCount, m.nextError
}

func TestNewTapeFromReader(t *testing.T) {
	testError := fmt.Errorf("test")

	defer func() {
		err := recover()
		if err != testError {
			t.Errorf("expected NewTapeFromReader to panic with test error, got %q", err)
		}
	}()

	reader := &MockReader{
		[]byte{},
		0,
		testError,
	}

	NewTapeFromReader(reader)

	t.Errorf("expected NewTapeFromReader to panic, but it succeeded")
}

func TestTapePositionControl(t *testing.T) {
	reader := NewTapeFromString("abc")

	if reader.Position() != 0 {
		t.Errorf("Expected reader to be initialized at position 0")
	}

	if reader.Advance() != 1 {
		t.Errorf("Expected reader.Advance() to return the next position")
	}
	if reader.Position() != 1 {
		t.Errorf("Expected reader.Advance() to adavance the position")
	}

	if reader.Backup() != 0 {
		t.Errorf("Expected reader.Backup() to return the previous position")
	}
	if reader.Position() != 0 {
		t.Errorf("Expected reader.Backup() to backup the position")
	}

	if reader.Backup() != -1 {
		t.Errorf("Expected reader.Backup() to allow backing up before 0")
	}

	reader.Advance()
	reader.Advance()
	reader.Advance()

	if reader.Advance() != 3 {
		t.Errorf("Expected reader.Advance() to allow advancing past length")
	}
}

func TestTapePeek(t *testing.T) {
	reader := NewTapeFromString("abc")
	a, _ := utf8.DecodeRuneInString("a")
	b, _ := utf8.DecodeRuneInString("b")
	c, _ := utf8.DecodeRuneInString("c")

	if reader.Peek() != a {
		t.Errorf("Expected reader.Peek() to return first rune after initialization")
	}

	reader.Advance()
	if reader.Peek() != b {
		t.Errorf("Expected reader.Peek() to return second rune after advancing once")
	}

	reader.Backup()
	reader.Backup()
	if reader.Peek() != EOFRune {
		t.Errorf("Expected reader.Peek() to return EOFRune when position is negative")
	}

	reader.Advance()
	reader.Advance()
	reader.Advance()
	if reader.Peek() != c {
		t.Errorf("Expected reader.Peek() to return last rune when position is at the last rune")
	}

	reader.Advance()
	if reader.Peek() != EOFRune {
		t.Errorf("Expected reader.Peek() to return EOFRune when position is after the end")
	}
}

func TestTapeNext(t *testing.T) {
	reader := NewTapeFromString("abc")
	a, _ := utf8.DecodeRuneInString("a")

	if reader.Next() != a {
		t.Errorf("Expected reader.Next() to return next rune")
	}

	if reader.Position() != 1 {
		t.Errorf("Expected reader.Next() to advance position")
	}

	reader.Next()
	reader.Next()

	if reader.Next() != EOFRune {
		t.Errorf("Expected reader.Next() to return eof when reading after data")
	}
}

func TestTapeRewind(t *testing.T) {
	reader := NewTapeFromString("abc")
	a, _ := utf8.DecodeRuneInString("a")

	reader.Next()

	if reader.Rewind() != a {
		t.Errorf("Expected reader.Rewind() to return new peek")
	}

	if reader.Position() != 0 {
		t.Errorf("Expected reader.Rewind() to backup position")
	}

	if reader.Rewind() != EOFRune {
		t.Errorf("Expected reader.Rewind() to return eof when reading before data")
	}
}

func TestReadSlice(t *testing.T) {
	tape := NewTapeFromString("abc")
	var slice []rune

	if slice = tape.ReadSlice(5, 3); len(slice) != 0 {
		t.Errorf("expected ReadSlice() to return an empty slice when start is after the end")
	}
}
