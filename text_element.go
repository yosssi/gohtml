package gohtml

import "bytes"

// A textElement represents a text element of an HTML document.
type textElement struct {
	text string
}

// write writes a text to the buffer.
func (e *textElement) write(bf *bytes.Buffer) {
	bf.WriteString(e.text)
}
