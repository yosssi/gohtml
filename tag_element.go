package gohtml

import "bytes"

// A tagElement represents a tag element of an HTML document.
type tagElement struct {
	tagName     string
	startTagRaw string
	endTagRaw   string
	children    []element
}

// Condense any tag with no child tags (only text or nothing) onto a single line
var Condense bool

// write writes a tag to the buffer.
func (e *tagElement) write(bf *bytes.Buffer, indent int) {
	if Condense {
		l := len(e.children)
		if l == 0 {
			writeLine(bf, indent, e.startTagRaw, e.endTagRaw)
			return
		} else if l == 1 && e.endTagRaw != "" {
			if c, ok := e.children[0].(*textElement); ok {
				writeLine(bf, indent, e.startTagRaw, c.text, e.endTagRaw)
				return
			}
		}
	}

	writeLine(bf, indent, e.startTagRaw)
	for _, c := range e.children {
		var childIndent int
		if e.endTagRaw != "" {
			childIndent = indent + 1
		} else {
			childIndent = indent
		}
		c.write(bf, childIndent)
	}
	if e.endTagRaw != "" {
		writeLine(bf, indent, e.endTagRaw)
	}
}

// appendChild append an element to the element's children.
func (e *tagElement) appendChild(child element) {
	e.children = append(e.children, child)
}
