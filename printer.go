package main

import (
	"io"
	"image"
	"image/gif"
	"image/color"
	"golang.org/x/image/draw"
)

type printer interface {
	print(v []bool)
	close() error
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

func (*logPrinter) close() error { return nil }

type imagePrinter struct {
	io.Writer
	*image.Gray
	size int
	generation int
}
//TODO configure square size in pixels as well?
func newImagePrinter(width, generations, size int, out io.Writer) printer {
	return &imagePrinter{
		Writer: out,
		Gray: image.NewGray(image.Rect(0, 0, width*size, generations*size)),
		size: size,
	}
}

func (p *imagePrinter) print(v []bool) {
	for i, b := range v {
		var c color.Color
		if b {
			c = color.Black
		} else {
			c = color.White
		}
		x := i * p.size
		y := p.generation * p.size
		rect := image.Rect(x, y, x + p.size, y + p.size)
		draw.Draw(p.Gray, rect, &image.Uniform{c}, image.ZP, draw.Src)
	}
	p.generation++
}

func (p *imagePrinter) close() error {
	return gif.Encode(p.Writer, p.Gray, &gif.Options{NumColors:2})
}

type svgPrinter struct {
	io.Writer
}

func (p *svgPrinter) print(v []bool) {
	//TODO print
}
