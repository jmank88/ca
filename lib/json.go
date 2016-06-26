package ca

import (
	"io"
)

func init() {
	Register("json", NewPrinter(func(cells, generations, size int, out io.Writer) Printer {
		return NewJsonPrinter(out)
	}))
}

type JsonPrinter struct {
	io.Writer
	first bool
}

func NewJsonPrinter(w io.Writer) Printer {
	io.WriteString(w, "[")
	return &JsonPrinter{Writer: w, first: true}
}

func (p *JsonPrinter) Print(v []bool) {
	if !p.first {
		io.WriteString(p, ",\n")
	}
	p.first = false
	io.WriteString(p, "[")
	for i, b := range v {
		if b {
			io.WriteString(p, "true")
		} else {
			io.WriteString(p, "false")
		}
		if i < len(v)-1 {
			io.WriteString(p, ",")
		}
	}
	io.WriteString(p, "]")
}

func (p *JsonPrinter) Close() error {
	io.WriteString(p, "]")
	return nil
}