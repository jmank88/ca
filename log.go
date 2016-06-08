package main

import "io"

const (
	blackSquare = "\u25FC"
	whiteSquare = "\u25FD"
)

// A logPrinter is a printer which logs automata cells.
type logPrinter struct {
	io.Writer
}

func (p *logPrinter) print(v []bool) {
	for _, b := range v {
		if b {
			io.WriteString(p, blackSquare)
		} else {
			io.WriteString(p, whiteSquare)
		}
	}
	io.WriteString(p, "\n")
}

func (*logPrinter) close() error { return nil }
