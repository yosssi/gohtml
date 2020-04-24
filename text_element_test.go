package gohtml

import (
	"bytes"
	"testing"
)

func TestTextElementWrite(t *testing.T) {
	textElem := &textElement{text: "Test text"}
	bf := &formattedBuffer{buffer: &bytes.Buffer{}}
	textElem.write(bf, true)
	actual := bf.buffer.String()
	expected := "Test text"
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}
