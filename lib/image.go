package ca

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func init() {
	Register("gif", NewPrinter(func(cells, generations, size int, out io.Writer) Printer {
		return NewImagePrinter(cells, generations, size, out, writeGIF)
	}))
	Register("png", NewPrinter(func(cells, generations, size int, out io.Writer) Printer {
		return NewImagePrinter(cells, generations, size, out, writePNG)
	}))
	jpg := NewPrinter(func(cells, generations, size int, out io.Writer) Printer {
		return NewImagePrinter(cells, generations, size, out, writeJPG)
	})
	Register("jpg", jpg)
	Register("jpeg", jpg)
}

// The newImagePrinter function returns a new image printer which writes to out.
func NewImagePrinter(cells, generations, size int, out io.Writer, onClose func(*ImageDrawer) error) Printer {
	return &DrawPrinter{
		Drawer: &ImageDrawer{
			Writer:  out,
			Gray:    image.NewGray(image.Rect(0, 0, cells*size, generations*size)),
			onClose: onClose,
		},
		size: size,
	}
}

// An ImageDrawer is a drawer which draws to a grayscale gif.
type ImageDrawer struct {
	io.Writer
	*image.Gray
	onClose func(*ImageDrawer) error
}

func (p *ImageDrawer) Draw(x0, y0, x1, y1 int, b bool) {
	var c color.Color
	if b {
		c = color.Black
	} else {
		c = color.White
	}
	rect := image.Rect(x0, y0, x1, y1)
	draw.Draw(p.Gray, rect, &image.Uniform{c}, image.ZP, draw.Src)
}

func (p *ImageDrawer) Close() error {
	return p.onClose(p)
}

//TODO no colors (works via gopherjs...)
func writeGIF(p *ImageDrawer) error {
	return gif.Encode(p.Writer, p.Gray, &gif.Options{NumColors: 2, Quantizer: &bwQuantizer{}})
}

type bwQuantizer struct{}

func (q *bwQuantizer) Quantize(p color.Palette, m image.Image) color.Palette {
	return color.Palette{color.Black, color.White}
}

func writePNG(p *ImageDrawer) error {
	return png.Encode(p.Writer, p.Gray)
}

func writeJPG(p *ImageDrawer) error {
	return jpeg.Encode(p.Writer, p.Gray, &jpeg.Options{Quality: 25})
}
