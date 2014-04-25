# GoHTML - HTML formatter for Go

[![Build Status](http://128.199.249.74/github.com/yosssi/gohtml/status.png?branch=master)](http://128.199.249.74/github.com/yosssi/gohtml)
[![GoDoc](https://godoc.org/github.com/yosssi/gohtml?status.png)](https://godoc.org/github.com/yosssi/gohtml)

GoHTML is an HTML formatter for [Go](http://golang.org/). You can format HTML source codes by using this package.

## Example

Example Go source code:

```go
package main

import (
	"fmt"

	"github.com/yosssi/gohtml"
)

func main() {
	h := `<!DOCTYPE html><html><head><title>This is a title.</title><script type="text/javascript">
alert('aaa');
if (0 < 1) {
	alert('bbb');
}
</script><style type="text/css">
body {font-size: 14px;}
h1 {
	font-size: 16px;
	font-weight: bold;
}
</style></head><body><form><input type="name"><p>AAA<br>BBB></p></form><!-- This is a comment. --></body></html>`
	fmt.Println(gohtml.Format(h))
}
```

Output:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>
      This is a title.
    </title>
    <script type="text/javascript">
      alert('aaa');
      if (0 < 1) {
      	alert('bbb');
      }
    </script>
    <style type="text/css">
      body {font-size: 14px;}
      h1 {
      	font-size: 16px;
      	font-weight: bold;
      }
    </style>
  </head>
  <body>
    <form>
      <input type="name">
      <p>
        AAA
        <br>
        BBB>
      </p>
    </form>
    <!-- This is a comment. -->
  </body>
</html>
```

## Docs

* [GoDoc](https://godoc.org/github.com/yosssi/gohtml)
