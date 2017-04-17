package ca

// The Drawer interface is an abstraction for drawing binary automata cells to an x/y grid.
type Drawer interface {
	Draw(x0, y0, x1, y1 int, b bool)
	Close() error
}

type DrawPrinter struct {
	Drawer
	// The height and width of each cell.
	size       int
	generation int
}

func (d *DrawPrinter) Print(v []bool) {
	for i, b := range v {
		x := i * d.size
		y := d.generation * d.size
		d.Draw(x, y, x+d.size, y+d.size, b)
	}
	d.generation++
}

func (d *DrawPrinter) Close() error {
	return d.Drawer.Close()
}
