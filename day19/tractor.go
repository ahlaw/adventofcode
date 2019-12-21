package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day18/intcode"
)

func check(ops []int64, x, y int) bool {
	prog := make([]int64, len(ops))
	copy(prog, ops)
	in := make(chan int)
	out := make(chan int)
	defer close(in)
	go intcode.RunProgram(prog, in, out)
	in <- x
	in <- y
	return <-out == 1
}

func part1(ops []int64) int {
	width := 50
	count := 0
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			if check(ops, x, y) {
				count++
			}
		}
	}
	return count
}

func part2(ops []int64) int {
	x := 0
	y := 0
	for {
		if check(ops, x, y+99) {
			if check(ops, x+99, y) {
				return x*10000 + y
			}
			y++
		} else {
			x++
		}
	}
}

func main() {
	ops := intcode.ReadProgram()
	fmt.Print("Part1: ")
	fmt.Println(part1(ops))
	fmt.Print("Part2: ")
	fmt.Println(part2(ops))
}
