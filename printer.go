package main

import (
	"io"
	"image"
)

type printer interface {
	print(v []bool)
}

const (
	blackSquare = "\u25FC"
	whiteSquare = "\u25FD"
)

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

type imagePrinter struct {
	image.Image
}

func (p *imagePrinter) print(v []bool) {
	//TODO print
}

type svgPrinter struct {
	io.Writer
}

func (p *svgPrinter) print(v []bool) {
	//TODO print
}
