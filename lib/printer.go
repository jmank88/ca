package ca

import "io"

// A Printer prints automata generations.
type Printer interface {
	Print(v []bool)
	Close() error
}

//TODO doc
type NewPrinter func(cells, generations, size int, out io.Writer) Printer
