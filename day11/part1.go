package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day11/intcode"
	"image"
)

func main() {
	ops := intcode.ReadProgram()
	in := make(chan int, 1)
	out := make(chan int)
	go intcode.RunProgram(ops, in, out)

	pos := image.Point{}
	dir := 0
	hull := make(map[image.Point]int)

	for {
		in <- hull[pos]
		hull[pos] = <-out
		turn, ok := <-out
		if !ok {
			break
		}
		dir = (dir + []int{-1, 1}[turn] + 4) % 4
		pos = pos.Add([]image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}[dir])
	}
	fmt.Println(len(hull))
}
