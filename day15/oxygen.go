package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day15/intcode"
	"image"
)

const (
	NORTH int = iota + 1
	SOUTH
	WEST
	EAST
)

var directions = map[int]image.Point{
	NORTH: {0, 1},
	SOUTH: {0, -1},
	WEST:  {-1, 0},
	EAST:  {1, 0},
}

type droid struct {
	in  chan<- int
	out <-chan int
}

func (d *droid) step(dir int) int {
	d.in <- dir
	return <-d.out
}

func opposite(dir int) int {
	switch dir {
	case NORTH:
		return SOUTH
	case SOUTH:
		return NORTH
	case WEST:
		return EAST
	case EAST:
		return WEST
	}
	return 0
}

func dfs(grid map[image.Point]int, currPos image.Point, d droid) {
	for dir, move := range directions {
		newPos := currPos.Add(move)
		if _, ok := grid[newPos]; ok {
			continue
		}
		status := d.step(dir)
		if grid[newPos] = status; status == 0 {
			continue
		}
		dfs(grid, newPos, d)
		d.step(opposite(dir))
	}
}

func bfs(grid map[image.Point]int, startPos image.Point) map[image.Point]int {
	dist := map[image.Point]int{startPos: 0}
	queue := []image.Point{startPos}
	for len(queue) > 0 {
		currPos := queue[0]
		queue = queue[1:]
		for _, move := range directions {
			newPos := currPos.Add(move)
			if _, ok := dist[newPos]; ok {
				continue
			}
			if grid[newPos] > 0 {
				dist[newPos] = dist[currPos] + 1
				queue = append(queue, newPos)
			}
		}
	}
	return dist
}

func locateOxygen(grid map[image.Point]int) image.Point {
	for loc, status := range grid {
		if status == 2 {
			return loc
		}
	}
	return image.Point{}
}

func largestDistance(distances map[image.Point]int) int {
	var max int
	for _, dist := range distances {
		if dist > max {
			max = dist
		}
	}
	return max
}

func main() {
	ops := intcode.ReadProgram()
	in := make(chan int)
	out := make(chan int)
	d := droid{in, out}
	go intcode.RunProgram(ops, in, out)
	start := image.Point{0, 0}
	grid := map[image.Point]int{start: 1}
	dfs(grid, start, d)
	oxygenPoint := locateOxygen(grid)
	distances := bfs(grid, oxygenPoint)
	fmt.Print("Part 1: ")
	fmt.Println(distances[start])
	fmt.Print("Part 2: ")
	fmt.Println(largestDistance(distances))
}
