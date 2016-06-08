package main

import (
	"flag"
	"log"
	"os"
	"crypto/rand"
	"math/big"
	"path/filepath"
)

const (
	width = 50
	generations = 50
)

var (
	ruleF = flag.Int("r", 110, "rule (0-255)")
	randF = flag.Bool("rand", false, "randomized initial state")
	fileF = flag.String("file", "", "output filename; recognized extensions: txt, svg, gif, json")
)

func main() {
	flag.Parse()

	if *ruleF > 255 || *ruleF < 0 {
		log.Fatalf("invalid rule (%d); must be between 0-255\n", *ruleF)
	}
	r := rule(*ruleF)

	var p printer

	if *fileF == "" {
		p = &logPrinter{os.Stdout}
	} else {
		switch filepath.Ext(*fileF) {
		case ".txt":
			if f, err := os.Create(*fileF); err != nil {
				log.Fatal(err)
			} else {
				defer f.Close()
				p = &logPrinter{f}
			}
		case ".svg":
			if f, err := os.Create(*fileF); err != nil {
				log.Fatal(err)
			} else {
				defer f.Close()
				p = newSvgPrinter(width, generations, 10, f)
			}
		case ".gif":
			if f, err := os.Create(*fileF); err != nil {
				log.Fatal(err)
			} else {
				defer f.Close()
				p = newImagePrinter(width, generations, 10, f, GIF)
			}
		case ".png":
			if f, err := os.Create(*fileF); err != nil {
				log.Fatal(err)
			} else {
				defer f.Close()
				p = newImagePrinter(width, generations, 10, f, PNG)
			}
		case ".jpeg",".jpg":
			if f, err := os.Create(*fileF); err != nil {
				log.Fatal(err)
			} else {
				defer f.Close()
				p = newImagePrinter(width, generations, 10, f, JPEG)
			}
		case ".json":
			if f, err := os.Create(*fileF); err != nil {
				log.Fatal(err)
			} else {
				defer f.Close()
				p = newJsonPrinter(f)
			}
		default:
			if f, err := os.Create(*fileF + ".txt"); err != nil {
				log.Fatal(err)
			} else {
				defer f.Close()
				p = &logPrinter{f}
			}
		}
	}

	var last, next [width]bool
	initialState(last[:], *randF)
	p.print(last[:])
	for i := 0; i < generations; i++ {
		for j := 0; j < width; j++ {
			var l int8
			if last[j] {
				l |= 1 << 1
			}
			if j > 0 && last[j-1] {
				l |= 1 << 2
			}
			if j < width - 1 && last[j+1] {
				l |= 1
			}
			next[j] = r.apply(l)
		}
		next, last = last, next
		p.print(last[:])
	}
	if err := p.close(); err != nil {
		log.Fatal(err)
	}
}

var maxstate = new(big.Int).SetInt64(2)

func initialState(v []bool, randomize bool) {
	if !randomize {
		v[width/2] = true
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
func (r *rule) apply(last int8) bool {
	return (1 << byte(last)) & *r > 0
}