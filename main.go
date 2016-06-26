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
	var file string

	flag.IntVar(&config.Rule, "r", 110, "rule (0-255)")
	flag.BoolVar(&config.Random, "rand", false, "randomized initial state")
	flag.IntVar(&config.Cells, "cells", 50, "number of cells")
	flag.IntVar(&config.Generations, "gens", 50, "generations")
	flag.StringVar(&file, "file", "", "output filename")
	flag.StringVar(&config.Format, "format", "", "output format; override file extension; one of: txt, svg, gif, json, png, jpg, jpeg")

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