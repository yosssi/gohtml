package gohtml

// Format parses the input HTML string, formats it and returns the result.
func Format(s string) string {
	return parse(s).html()
}
