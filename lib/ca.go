package ca

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

var types = map[string]NewPrinter{}

func Register(key string, np NewPrinter) {
	types[key] = np
}

var Default = Config{
	Cells: 50,
	Generations: 50,
	Format: "txt",
	Random: false,
	Rule: 30,
}

type Config struct {
	Cells       int
	Generations int
	Format      string
	Random      bool
	Rule        int
}

func (c Config) Print(w io.Writer) error {

	if c.Cells < 1 {
		c.Cells = Default.Cells
	}

	if c.Generations < 1 {
		c.Generations = Default.Generations
	}

	if c.Rule > 255 || c.Rule < 0 {
		c.Rule = Default.Rule
	}
	r := rule(c.Rule)

	np, ok := types[c.Format]
	if !ok {
		return fmt.Errorf("invalid type: %s", c.Format)
	}

	p := np(c.Cells, c.Generations, c.Cells, w)

	last := make([]bool, c.Cells, c.Cells)
	next := make([]bool, c.Cells, c.Cells)
	initialState(last, c.Random)
	p.Print(last)
	for i := 0; i < c.Generations; i++ {
		for j := 0; j < c.Cells; j++ {
			var l int8
			if last[j] {
				l |= 1 << 1
			}
			if j > 0 && last[j-1] {
				l |= 1 << 2
			}
				if j < c.Cells - 1 && last[j+1] {
				l |= 1
			}
			next[j] = r.apply(l)
		}
		next, last = last, next
		p.Print(last)
	}

	return p.Close()
}

var maxstate = new(big.Int).SetInt64(2)

func initialState(v []bool, randomize bool) {
	if !randomize {
		v[len(v)/2] = true
		return
	}

	for i, _ := range v {
		if r, err := rand.Int(rand.Reader, maxstate); err != nil {
			panic(err)
		} else {
			v[i] = r.Int64() == 0
		}
	}
}

type rule int8

// Note: last must be between 0-8
func (r rule) apply(last int8) bool {
	return (1 << byte(last)) & r > 0
}