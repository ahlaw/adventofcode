package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day11/intcode"
	"image"
)

func drawHull(hull map[image.Point]int) {
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for pos := range hull {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			col := hull[image.Point{x, y}]
			if col == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	ops := intcode.ReadProgram()
	in := make(chan int, 1)
	out := make(chan int)
	go intcode.RunProgram(ops, in, out)

	pos := image.Point{}
	dir := 0
	hull := make(map[image.Point]int)

	in <- 1
	for {
		hull[pos] = <-out
		turn, ok := <-out
		if !ok {
			break
		}
		dir = (dir + []int{-1, 1}[turn] + 4) % 4
		pos = pos.Add([]image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}[dir])
		in <- hull[pos]
	}
	drawHull(hull)
}
