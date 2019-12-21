package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day17/intcode"
	"image"
)

var moves = [4]image.Point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func readMap(out <-chan int) (map[image.Point]rune, int, int) {
	grid := make(map[image.Point]rune)
	var x, y int
	var width, height int
	for c := range out {
		if c != 10 {
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

func findIntersections(grid map[image.Point]rune, width, height int) []image.Point {
	intersections := make([]image.Point, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pos := image.Point{x, y}
			if grid[pos] == '#' {
				isIntersection := true
				for _, move := range moves {
					adj := pos.Add(move)
					if grid[adj] != '#' {
						isIntersection = false
					}
				}
				if isIntersection {
					intersections = append(intersections, pos)
				}
			}
		}
	}
	return intersections
}

func sumAlignment(intersections []image.Point) int {
	var sum int
	for _, intersection := range intersections {
		sum += intersection.X * intersection.Y
	}
	return sum
}

func drawMap(grid map[image.Point]rune, width, height int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pos := image.Point{x, y}
			fmt.Printf("%q", grid[pos])
		}
		fmt.Println()
	}
}

func main() {
	ops := intcode.ReadProgram()
	in := make(chan int)
	out := make(chan int)
	go intcode.RunProgram(ops, in, out)
	grid, width, height := readMap(out)
	intersections := findIntersections(grid, width, height)
	fmt.Println(sumAlignment(intersections))
	//drawMap(grid, width, height)
}
