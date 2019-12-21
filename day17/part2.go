package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day17/intcode"
	"image"
)

var moves = [4]image.Point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

var moveInput = [4]string{
	"A,B,A,C,B,C,B,C,A,C",
	"L,10,R,12,R,12",
	"R,6,R,10,L,10",
	"R,10,L,10,L,12,R,6",
}

func readMap(out <-chan int) (map[image.Point]rune, int, int) {
	grid := make(map[image.Point]rune)
	var x, y int
	var width, height int
	for c := range out {
		if c != '\n' {
			grid[image.Point{x, y}] = rune(c)
			x++
		} else {
			if x == 0 {
				height = y
				return grid, width, height
			}
			y++
			width = x
			x = 0
		}
	}
	height = y
	return grid, width, height
}

func drawMap(grid map[image.Point]rune, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := image.Point{x, y}
			fmt.Printf("%q", grid[pos])
		}
		fmt.Println()
	}
}

func readPrompt(out <-chan int) {
	for {
		c := <-out
		if c == '\n' {
			break
		}
	}
}

func writeMove(in chan<- int, line string) {
	for _, c := range line {
		in <- int(c)
	}
	in <- int('\n')
}

func computeDust() int {
	ops := intcode.ReadProgram()
	ops[0] = 2
	in := make(chan int)
	out := make(chan int)
	go intcode.RunProgram(ops, in, out)
	readMap(out)
	for _, line := range moveInput {
		readPrompt(out)
		writeMove(in, line)
	}
	readPrompt(out)
	writeMove(in, "n")
	var dust int
	for output := range out {
		dust = output
	}
	return dust
}

func main() {
	/*
		ops := intcode.ReadProgram()
		in := make(chan int)
		out := make(chan int)
		go intcode.RunProgram(ops, in, out)
		grid, width, height := readMap(out)
		drawMap(grid, width, height)
	*/
	fmt.Println(computeDust())
}
