package gohtml

import "bytes"

// A tagElement represents a tag element of an HTML document.
type tagElement struct {
	tagName     string
	startTagRaw string
	endTagRaw   string
	children    []element
}

// write writes a tag to the buffer.
func (e *tagElement) write(bf *bytes.Buffer) {
	bf.WriteString(e.startTagRaw)
	for _, c := range e.children {
		c.write(bf)
	}
	bf.WriteString(e.endTagRaw)
}

// appendChild append an element to the element's children.
func (e *tagElement) appendChild(child element) {
	e.children = append(e.children, child)
}
