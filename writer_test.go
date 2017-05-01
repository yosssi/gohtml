package gohtml

import (
	"bytes"
	"testing"
)

func TestWriterSetLastElement(t *testing.T) {
	buf := &bytes.Buffer{}
	wr := NewWriter(buf)
	wr.SetLastElement("test")
	if wr.lastElement != "test" {
		t.Errorf("Invalid lastElement. [expected: %s][actual: %s]", "test", wr.lastElement)
	}
}

func TestWriterWrite(t *testing.T) {
	buf := &bytes.Buffer{}
	wr := NewWriter(buf)
	n, err := wr.Write([]byte("<html><head><title>This is a title.</title></head><body><p>test</p></body></html>"))
	if err != nil {
		t.Errorf("An error occurred. [error: %s]", err.Error())
	}
	expected := 129
	if n != expected {
		t.Errorf("Invalid return value. [expected: %d][actual: %d]", expected, n)
	}

	buf = &bytes.Buffer{}
	wr = NewWriter(buf)
	n, err = wr.Write([]byte(""))
	if err != nil {
		t.Errorf("An error occurred. [error: %s]", err.Error())
	}
	expected = 0
	if n != expected {
		t.Errorf("Invalid return value. [expected: %d][actual: %d]", expected, n)
	}
}

func TestNewWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	wr := NewWriter(buf)
	if wr.writer != buf || wr.lastElement != defaultLastElement || wr.bf.Len() != 0 {
		t.Errorf("Invalid Writer. [expected: %+v][actual: %+v]", &Writer{writer: buf, lastElement: defaultLastElement, bf: &bytes.Buffer{}}, wr)
	}
}
