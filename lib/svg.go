package ca

import (
	"fmt"
	"io"

	"github.com/ajstarks/svgo"
)

func init() {
	svg := func(cells, generations, size int, out io.Writer) Printer {
		return NewSvgPrinter(cells, generations, size, out)
	}
	Register("svg", svg)
}

// An SvgDrawer is a drawer
type SvgDrawer struct {
	*svg.SVG
}

func NewSvgPrinter(cells, generations, size int, out io.Writer) Printer {
	s := svg.New(out)
	x := cells * size
	y := generations * size
	s.Startraw(fmt.Sprintf(`viewBox="0 0 %d %d"`, x, y))
	s.Rect(0, 0, x, y, "fill:#ffffff;")
	return &DrawPrinter{
		Drawer: &SvgDrawer{
			SVG: s,
		},
		size: size,
	}
}

func (s *SvgDrawer) Draw(x0, y0, x1, y1 int, b bool) {
	if b {
		s.Rect(x0, y0, x1-x0, y1-y0, "fill:#000000;")
	}
}

func (s *SvgDrawer) Close() error {
	s.End()
	return nil
}
