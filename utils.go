package gohtml

import (
	"bytes"
	"strings"
)

// writeLine writes an HTML line to the buffer.
func writeLine(bf *bytes.Buffer, s string, indent int) {
	writeLineFeed(bf)
	writeIndent(bf, indent)
	bf.WriteString(s)
}

// writeLineFeed writes a line feed to the buffer.
func writeLineFeed(bf *bytes.Buffer) {
	if bf.Len() > 0 {
		bf.WriteString("\n")
	}
}

// writeIndent writes indents to the buffer.
func writeIndent(bf *bytes.Buffer, indent int) {
	bf.WriteString(strings.Repeat(DefaultIndentString, indent))
}
