package ca

import "io"

func init() {
	log := NewPrinter(func(cells, generations, size int, out io.Writer) Printer {
		return NewLogPrinter(cells, generations, size, out)
	})
	Register("", log)
	Register("log", log)
	Register("txt", log)
}

const (
	whiteSquare = "\u25FD"
	blackSquare = "\u25FE"
)

func NewLogPrinter(cells, generations, size int, out io.Writer) Printer {
	return &logPrinter{out}
}

// A logPrinter is a printer which logs automata cells.
type logPrinter struct {
	io.Writer
}

func (p *logPrinter) Print(v []bool) {
	for _, b := range v {
		if b {
			io.WriteString(p, blackSquare)
		} else {
			io.WriteString(p, whiteSquare)
		}
	}
	io.WriteString(p, "\n")
}

func (*logPrinter) Close() error { return nil }
