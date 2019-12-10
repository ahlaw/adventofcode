package main

import (
	"github.com/ahlaw/adventofcode/day09/intcode"
)

func main() {
	ops := intcode.ReadProgram()
	intcode.RunProgram(ops)
}
