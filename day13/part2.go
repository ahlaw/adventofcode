package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day13/intcode"
)

func main() {
	ops := intcode.ReadProgram()
	ops[0] = 2
	in := make(chan int, 1)
	out := make(chan int)
	go intcode.RunProgram(ops, in, out)

	score := 0
	paddleX := 0
	for x := range out {
		y := <-out
		id := <-out
		if x == -1 && y == 0 {
			score = id
		} else if id == 3 {
			paddleX = x
		} else if id == 4 {
			switch {
			case paddleX > x:
				in <- -1
			case paddleX < x:
				in <- 1
			default:
				in <- 0
			}
		}
	}
	fmt.Println(score)
}
