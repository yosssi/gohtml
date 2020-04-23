package gohtml

// An element represents an HTML element.
type element interface {
	write(*formattedBuffer, bool)
}
