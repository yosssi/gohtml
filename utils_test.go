package gohtml

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteLineFeed(t *testing.T) {
	bf := &formattedBuffer{
		buffer: &bytes.Buffer{},
	}
	bf.writeToken("test  \n ", formatterTokenType_Text)
	bf.writeLineFeed()
	actual := bf.buffer.String()
	expected := "test  \n\n"
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestWriteIndent(t *testing.T) {
	bf := &formattedBuffer{
		buffer:       &bytes.Buffer{},
		indentString: "  ",
		indentLevel:  3,
	}
	bf.writeIndent()
	actual := bf.buffer.String()
	expected := "      "
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestWriteTokenSpacing(t *testing.T) {
	bf := &formattedBuffer{
		buffer: &bytes.Buffer{},
	}

	bf.writeToken("text1", formatterTokenType_Text)
	bf.writeToken("text2", formatterTokenType_Text)
	bf.writeToken("text3", formatterTokenType_Text)
	bf.writeLineFeed()
	bf.writeToken("text1", formatterTokenType_Text)
	bf.writeToken("<tag1>", formatterTokenType_Tag)
	bf.writeToken("<tag2>", formatterTokenType_Tag)
	bf.writeToken("text2", formatterTokenType_Text)
	bf.writeLineFeed()
	bf.writeToken("<tag1>", formatterTokenType_Tag)
	bf.writeToken("text1", formatterTokenType_Text)
	bf.writeToken("text2", formatterTokenType_Text)
	bf.writeToken("<tag2>", formatterTokenType_Tag)
	bf.writeLineFeed()
	bf.writeToken("<tag1>", formatterTokenType_Tag)
	bf.writeToken("<tag2>", formatterTokenType_Tag)
	bf.writeToken("<tag3>", formatterTokenType_Tag)

	actual := bf.buffer.String()
	expected := "text1 text2 text3\n" +
		"text1<tag1><tag2>text2\n" +
		"<tag1>text1 text2<tag2>\n" +
		"<tag1><tag2><tag3>"
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestWriteTokenWrapping(t *testing.T) {
	bf := &formattedBuffer{
		buffer:         &bytes.Buffer{},
		lineWrapColumn: 80,
	}

	for _, word := range strings.Split("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", " ") {
		bf.writeToken(word, formatterTokenType_Text)
	}

	actual := bf.buffer.String()
	expected := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor\n" +
		"incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis\n" +
		"nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.\n" +
		"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu\n" +
		"fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in\n" +
		"culpa qui officia deserunt mollit anim id est laborum."
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestWriteTokenWrappingWithSpillover(t *testing.T) {
	bf := &formattedBuffer{
		buffer:               &bytes.Buffer{},
		lineWrapColumn:       80,
		lineWrapMaxSpillover: 5,
	}

	for _, word := range strings.Split("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", " ") {
		bf.writeToken(word, formatterTokenType_Text)
	}

	actual := bf.buffer.String()
	expected := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor\n" +
		"incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud\n" +
		"exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute\n" +
		"irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat\n" +
		"nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa\n" +
		"qui officia deserunt mollit anim id est laborum."
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestUnifyLineFeed(t *testing.T) {
	actual := unifyLineFeed("\r\n\n\r")
	expected := "\n\n\n"
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}
