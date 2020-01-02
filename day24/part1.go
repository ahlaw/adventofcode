package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

const (
	BUGS  int = 1
	SPACE int = 0
)

func readGrid(path string) map[image.Point]int {
	f, _ := os.Open(path)
	defer f.Close()

	grid := make(map[image.Point]int)
	sc := bufio.NewScanner(f)
	y := 0
	for sc.Scan() {
		for x, c := range sc.Text() {
			if c == '#' {
				grid[image.Point{x, y}] = BUGS
			} else {
				grid[image.Point{x, y}] = SPACE
			}
		}
		y++
	}
	return grid
}

func timestep(currState map[image.Point]int) map[image.Point]int {
	newState := make(map[image.Point]int)
	for pos, tile := range currState {
		aliveNeighbors := 0
		for _, neighbor := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			aliveNeighbors += currState[pos.Add(neighbor)]
		}
		if tile == BUGS {
			if aliveNeighbors == 1 {
				newState[pos] = BUGS
			} else {
				newState[pos] = SPACE
			}
		} else {
			if aliveNeighbors == 1 || aliveNeighbors == 2 {
				newState[pos] = BUGS
			} else {
				newState[pos] = SPACE
			}
		}
	}
	return newState
}

func stepToRepeat(state map[image.Point]int) int {
	// each layout has unique biodiversity rating
	layouts := make(map[int]bool)
	for {
		state = timestep(state)
		rating := bioRating(state)
		if layouts[rating] {
			return rating
		}
		layouts[rating] = true
	}
	return -1
}

func bioRating(state map[image.Point]int) int {
	rating := 0
	for pos, tile := range state {
		if tile == BUGS {
			rating += 1 << (5*pos.Y + pos.X)
		}
	}
	return rating
}

func main() {
	grid := readGrid("./input.txt")
	fmt.Println(stepToRepeat(grid))
}
