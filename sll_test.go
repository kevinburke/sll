package main

import (
	"bytes"
	"testing"
)

func TestLinesStripped(t *testing.T) {
	b := new(bytes.Buffer)
	// Need to have at least 16 in the buffer to trigger NewBufferSize.
	input := `foo
012345678901234567890123
barbazbang
012345678901234567890123
012345678901234567890123
012345678901234567890123
blah`
	expected := `foo
barbazbang
blah
`
	r := bytes.NewReader([]byte(input))
	err := stripLongLines(16, r, b)
	if err != nil {
		t.Fatal(err)
	}
	s := b.String()
	if s != expected {
		t.Errorf("expected strip(input) to be %s, was %s", expected, s)
	}
}
