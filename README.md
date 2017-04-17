# Cellular Automata [![GoDoc](https://godoc.org/github.com/jmank88/ca?status.svg)](https://godoc.org/github.com/jmank88/ca) [![Go Report Card](https://goreportcard.com/badge/github.com/jmank88/ca)](https://goreportcard.com/report/github.com/jmank88/ca)

A program for generating cellular automata.

![Rule 30](example/30.png "Rule 30")

[View in Playground](https://jmank88.github.io/ca/playground/?rule=30)

## Usage

```
> ca --help
Usage of ca:
  -cells int
    	number of cells (default 50)
  -file string
    	output filename
  -format string
    	output format; override file extension; one of: txt, svg, gif, json, png, jpg, jpeg
  -gens int
    	generations (default 50)
  -r int
    	rule (0-255) (default 30)
  -rand
    	randomized initial state
```

See the go generate commands in [example.go](example.go).

## TODO

[ ] Input files (json, etc).

[ ] Fix gif output.
