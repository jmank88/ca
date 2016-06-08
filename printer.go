package main

// A printer prints automata generations.
type printer interface {
	print(v []bool)
	close() error
}
