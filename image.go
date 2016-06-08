package main

import (
	"io"
	"image"
	"image/color"
	"image/gif"
	"image/draw"
)

// An imageDrawer is a drawer which draws to a grayscale gif.
type imageDrawer struct {
	io.Writer
	*image.Gray
}

// The newImagePrinter function returns a new gif image printer which writes to out.
func newImagePrinter(width, generations, size int, out io.Writer) printer {
	return &drawPrinter{
		drawer: &imageDrawer{
			Writer: out,
			Gray: image.NewGray(image.Rect(0, 0, width * size, generations * size)),
		},
		size: size,
	}
}

func (p *imageDrawer) draw(x0, y0, x1, y1 int, b bool) {
	var c color.Color
	if b {
		c = color.Black
	} else {
		c = color.White
	}
	rect := image.Rect(x0, x1, y0, y1)
	draw.Draw(p.Gray, rect, &image.Uniform{c}, image.ZP, draw.Src)
}

func (p *imageDrawer) close() error {
	return gif.Encode(p.Writer, p.Gray, &gif.Options{NumColors:2})
}
