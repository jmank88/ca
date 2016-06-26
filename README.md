# Cellular Automata

A program for generating cellular automata.

![Rule 110](example/rand-110.png "Rule 110")
TODO link to webview for this and any others

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