package gohtml

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	h := `<html><head><title>This is a title</title><script>
	alert(1);
	alert(2);
</script><body><a href="http://google.com">To Google</a><p>AAA<br>BBB</p><input type="text"><br/><a>bbb<br>aaa</a><script type="text/javascript">if(0 < 1) {
	alert(111);
}</script></body></html>`

	htmlDoc := parse(h)
	fmt.Println(htmlDoc.html())
}
