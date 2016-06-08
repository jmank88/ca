package main

import (
	"io"
	"image"
	"image/color"
	"image/gif"
	"image/draw"
	"image/png"
	"image/jpeg"
)

// An imageDrawer is a drawer which draws to a grayscale gif.
type imageDrawer struct {
	io.Writer
	*image.Gray
	onClose func(*imageDrawer) error
}

type Type int

const (
	GIF Type = iota
	PNG
	JPEG
)

// The newImagePrinter function returns a new gif image printer which writes to out.
func newImagePrinter(width, generations, size int, out io.Writer, t Type) printer {
	var onClose func(*imageDrawer) error
	switch t {
	case GIF:
		onClose = writeGIF
	case PNG:
		onClose = writePNG
	case JPEG:
		onClose = writeJPG
	}
	return &drawPrinter{
		drawer: &imageDrawer{
			Writer: out,
			Gray: image.NewGray(image.Rect(0, 0, width * size, generations * size)),
			onClose: onClose,
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
	rect := image.Rect(x0, y0, x1, y1)
	draw.Draw(p.Gray, rect, &image.Uniform{c}, image.ZP, draw.Src)
}

func (p *imageDrawer) close() error {
	return p.onClose(p)
}

//TODO no colors
func writeGIF(p *imageDrawer) error {
	return gif.Encode(p.Writer, p.Gray, &gif.Options{NumColors:2, Quantizer: &bwQuantizer{}})
}

type bwQuantizer struct {}

func (q *bwQuantizer) Quantize(p color.Palette, m image.Image) color.Palette {
	return color.Palette{color.Black, color.White}
}

func writePNG(p *imageDrawer) error {
	return png.Encode(p.Writer, p.Gray)
}

func writeJPG(p *imageDrawer) error {
	return jpeg.Encode(p.Writer, p.Gray, &jpeg.Options{Quality:25})
}
