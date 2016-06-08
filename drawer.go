package main

// The drawer interface is an abstraction for drawing binary automata cells to an x/y grid.
type drawer interface {
	draw(x0, y0, x1, y1 int, b bool)
	close() error
}

type drawPrinter struct {
	drawer
	// The height and width of each cell.
	size int
	generation int
}

func (d *drawPrinter) print(v []bool) {
	for i, b := range v {
		x := i * d.size
		y := d.generation * d.size
		d.draw(x, y, x + d.size, y + d.size, b)
	}
	d.generation++
}

func (d *drawPrinter) close() error {
	return d.drawer.close()
}


