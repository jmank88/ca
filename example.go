//go:generate rm -f example/*
//go:generate ca -file example/18.txt -r 18
//go:generate ca -file example/rand-30.svg -rand -r 30 -cells 25 -gens 25
//go:generate ca -file example/110.json -r 110
//go:generate ca -file example/30.gif -r 30
//go:generate ca -file example/rand-30.png -rand -r 30 -cells 50 -gens 25
//go:generate ca -file example/30.png -r 30 -cells 30 -gens 25
//go:generate ca -file example/22.jpeg -r 22
//go:generate ca -file example/54.jpg -r 54
package main
