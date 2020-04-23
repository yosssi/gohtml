package gohtml

import (
	"regexp"
	"strings"
)

// A textElement represents a text element of an HTML document.
type textElement struct {
	text string
}

// write writes a text to the buffer.
func (e *textElement) write(bf *formattedBuffer, isPreviousNodeInline bool) {
	text := unifyLineFeed(e.text)

	if !isPreviousNodeInline {
		bf.writeLineFeed()
	}

	// Collapse leading and trailing spaces
	text = regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(text, " ")
	lines := strings.Split(text, "\n")
	for l, line := range lines {
		if l > 0 {
			bf.writeLineFeed()
		}
		for _, word := range strings.Split(line, " ") {
			bf.writeToken(word, formatterTokenType_Text)
		}
	}
}
