package main

import (
	"io"

	"github.com/ajstarks/svgo"
)

// An svgDrawer is a drawer which
type svgDrawer struct {
	*svg.SVG
}

func newSvgPrinter(width, generations, size int, out io.Writer) printer {
	s := svg.New(out)
	x := width * size
	y := generations * size
	s.Start(x, y)
	s.Rect(0, 0, x, y, "fill:#ffffff;")
	return &drawPrinter{
		drawer: &svgDrawer{
			SVG: s,
		},
		size: size,
	}
}

func (s *svgDrawer) draw(x0, y0, x1, y1 int, b bool) {
	if b {
		s.Rect(x0, y0, x1-x0, y1-y0, "fill:#000000;")
	}
}

func (s *svgDrawer) close() error {
	s.End()
	return nil
}