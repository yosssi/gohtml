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
func (e *tagElement) write(bf *formattedBuffer, isPreviousNodeInline bool) {
	if Condense && e.endTagRaw != "" {
		// Write the condensed output to a separate buffer, in case it doesn't work out
		condensedBuffer := *bf
		condensedBuffer.buffer = &bytes.Buffer{}

		if bf.buffer.Len() > 0 && !isPreviousNodeInline {
			condensedBuffer.writeLineFeed()
		}
		condensedBuffer.writeToken(e.startTagRaw, formatterTokenType_Tag)
		if !isPreviousNodeInline {
			condensedBuffer.indentLevel++
		}

		textOnly := true
		for _, child := range e.children {
			if _, ok := child.(*textElement); ok {
				child.write(&condensedBuffer, true)
			} else {
				textOnly = false
				break
			}
		}
		condensedBuffer.writeToken(e.endTagRaw, formatterTokenType_Tag)
		if !isPreviousNodeInline {
			condensedBuffer.indentLevel--
		}

		if textOnly && bytes.IndexAny(condensedBuffer.buffer.Bytes()[1:], "\n") == -1 {
			// If it was only text, and there were no newlines were in the buffer,
			// replace the original with the condensed version
			condensedBuffer.buffer = bytes.NewBuffer(bytes.Join([][]byte{
				bf.buffer.Bytes(), condensedBuffer.buffer.Bytes(),
			}, []byte{}))
			*bf = condensedBuffer

			return
		}
	}

	if bf.buffer.Len() > 0 {
		bf.writeLineFeed()
	}
	bf.writeToken(e.startTagRaw, formatterTokenType_Tag)
	if e.endTagRaw != "" {
		bf.indentLevel++
	}

	for _, child := range e.children {
		child.write(bf, false)
	}
	if e.endTagRaw != "" {
		bf.writeLineFeed()
		bf.indentLevel--
		bf.writeToken(e.endTagRaw, formatterTokenType_Tag)
	}
}

// appendChild append an element to the element's children.
func (e *tagElement) appendChild(child element) {
	e.children = append(e.children, child)
}
