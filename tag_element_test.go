package gohtml

import (
	"bytes"
	"testing"
)

func TestTagElementWrite(t *testing.T) {
	tagElem := &tagElement{tagName: "body", startTagRaw: "<body>", endTagRaw: "</body>"}
	child := &tagElement{tagName: "input", startTagRaw: "<input>"}
	grandChild := &textElement{text: "Test text"}
	child.appendChild(grandChild)
	tagElem.appendChild(child)
	bf := &formattedBuffer{buffer: &bytes.Buffer{}, indentString: defaultIndentString}
	tagElem.write(bf, true)
	actual := bf.buffer.String()
	expected := `<body>
  <input>
  Test text
</body>`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestTagElementAppendChild(t *testing.T) {
	tagElem := &tagElement{}
	child := &textElement{text: "TestText"}
	tagElem.appendChild(child)
	if len(tagElem.children) != 1 || tagElem.children[0] != child {
		t.Errorf("tagElement.children is invalid. [expected: %+v][actual: %+v]", []element{child}, tagElem.children)
	}
}
