package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jmank88/ca/lib"
)

func main() {
	var config ca.Config

	flag.IntVar(&config.Cells, "cells", ca.Default.Cells, "number of cells")
	flag.IntVar(&config.Generations, "gens", ca.Default.Generations, "generations")
	flag.StringVar(&config.Format, "format", ca.Default.Format, "output format; override file extension; one of: txt, svg, gif, json, png, jpg, jpeg")
	flag.BoolVar(&config.Random, "rand", ca.Default.Random, "randomized initial state")
	flag.IntVar(&config.Rule, "r", ca.Default.Rule, "rule (0-255)")

	var file string

	flag.StringVar(&file, "file", "", "output filename")

	flag.Parse()

	var out io.Writer

	if file == "" {
		out = os.Stdout
	} else {
		if f, err := os.Create(file); err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
			out = f
		}
	}

	if config.Format == "" {
		ext := filepath.Ext(file)
		if ext != "" {
			ext = ext[1:len(ext)]
		}
		config.Format = ext
	}

	if err := config.Print(out); err != nil {
		log.Fatal(err)
	}
}