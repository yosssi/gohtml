package gohtml

import (
	"bytes"
	"io"
)

type Writer struct {
	writer      io.Writer
	lastElement string
	bf          *bytes.Buffer
}

func (wr *Writer) SetLastElement(lastElement string) *Writer {
	wr.lastElement = lastElement
	return wr
}

func (wr *Writer) Write(p []byte) (n int, err error) {
	wr.bf.Write(p)
	if bytes.HasSuffix(p, []byte(wr.lastElement)) {
		return wr.writer.Write([]byte(Format(wr.bf.String()) + "\n"))
	}
	return 0, nil
}

func NewWriter(wr io.Writer) *Writer {
	return &Writer{writer: wr, lastElement: defaultLastElement, bf: &bytes.Buffer{}}
}
