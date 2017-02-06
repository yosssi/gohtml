package gohtml

import "bytes"

// An htmlDocument represents an HTML document.
type htmlDocument struct {
	elements []element
}

// html generates an HTML source code and returns it.
func (htmlDoc *htmlDocument) html() string {
	return string(htmlDoc.bytes())
}

// bytes reads from htmlDocument's internal array of elements and returns HTML source code
func (htmlDoc *htmlDocument) bytes() []byte {
	bf := &bytes.Buffer{}
	for _, e := range htmlDoc.elements {
		e.write(bf, startIndent)
	}
	return bf.Bytes()
}

// append appends an element to the htmlDocument.
func (htmlDoc *htmlDocument) append(e element) {
	htmlDoc.elements = append(htmlDoc.elements, e)
}
