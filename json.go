package main

import (
	"io"
)

type jsonPrinter struct {
	io.Writer
	first bool
}

func newJsonPrinter(w io.Writer) printer {
	io.WriteString(w, "[")
	return &jsonPrinter{Writer: w, first: true}
}

func (p *jsonPrinter) print(v []bool) {
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

func (p *jsonPrinter) close() error {
	io.WriteString(p, "]")
	return nil
}