package gohtml

import (
	"strings"
	"testing"
)

func TestHTMLDocumentHTML(t *testing.T) {
	s := `<!DOCTYPE html><html><head><title>This is a title.</title></head><body><p>Line1<br>Line2</p><br/></body></html><!-- aaa -->`
	htmlDoc := parse(strings.NewReader(s))

	actual := htmlDoc.html()
	expected := `<!DOCTYPE html>
<html>
  <head>
    <title>
      This is a title.
    </title>
  </head>
  <body>
    <p>
      Line1
      <br>
      Line2
    </p>
    <br/>
  </body>
</html>
<!-- aaa -->`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}

	// Try again to test idempotency
	htmlDoc = parse(strings.NewReader(actual))
	actual = htmlDoc.html()
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestHTMLDocumentAppend(t *testing.T) {
	htmlDoc := &htmlDocument{}
	textElem := &textElement{text: "TestText"}
	htmlDoc.append(textElem)
	if len(htmlDoc.elements) != 1 || htmlDoc.elements[0] != textElem {
		t.Errorf("htmlDocument.elements is invalid. [expected: %+v][actual: %+v]", []element{textElem}, htmlDoc.elements)
	}
}

func TestCondense(t *testing.T) {
	Condense = true
	defer func() {
		Condense = false
	}()
	s := `<!DOCTYPE html><html><head><title>This is a title.</title></head>` +
		`<body><p>` +
		`<strong><code><a>In</a></code>Line</strong>1<br>` +
		`Line2<br />` +
		`<em>Not<div>Inline</div></em>3` +
		`<strong>Un-<a href="Lorem ipsum dolor sit amet, consectetur adipiscing elit">inlined4</a></strong>` +
		`</p><p>A Single Line</p><br/>` +
		`</body></html><!-- aaa -->`
	htmlDoc := parse(strings.NewReader(s))
	actual := htmlDoc.html()
	expected := `<!DOCTYPE html>
<html>
  <head>
    <title>This is a title.</title>
  </head>
  <body>
    <p>
      <strong><code><a>In</a></code>Line</strong>1
      <br>
      Line2
      <br />
      <em>
        Not
        <div>Inline</div>
      </em>
      3
      <strong>
        Un-
        <a href="Lorem ipsum dolor sit amet, consectetur adipiscing elit">inlined4</a>
      </strong>
    </p>
    <p>A Single Line</p>
    <br/>
  </body>
</html>
<!-- aaa -->`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestHTMLTextWithNewline(t *testing.T) {
	s := `
<!DOCTYPE html><html><head></head><body>
<div>
  <span>
    I am content,

      <strong>spaced

        a bit weird.</strong>
  </span>
</div>
</body></html>
	`
	htmlDoc := parse(strings.NewReader(s))

	actual := htmlDoc.html()
	expected := `<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div>
      <span>
        I am content,
        <strong>
          spaced

          a bit weird.
        </strong>
      </span>
    </div>
  </body>
</html>`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}
