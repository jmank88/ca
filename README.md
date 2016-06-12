# Cellular Automata

A program for generating cellular automata.

![Rule 18](example/18.png "Rule 18")

## Usage

```
ca --help
Usage of ca:
  -file string
    	output filename; recognized extensions: txt, svg, gif, json
  -r int
    	rule (0-255) (default 110)
  -rand
    	randomized initial state
```

See the go generate commands in [example.go](example.go).

## TODO

[ ] Flags for width, generations, cell size

[ ] Input files (json, etc).

[ ] Fix gif output.