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

func readBase(path string) [5][5]int {
	f, _ := os.Open(path)
	defer f.Close()

	base := [5][5]int{}
	sc := bufio.NewScanner(f)
	y := 0
	for sc.Scan() {
		for x, c := range sc.Text() {
			if c == '#' {
				base[y][x] = BUGS
			} else {
				base[y][x] = SPACE
			}
		}
		y++
	}
	return base
}

func timestep(base [5][5]int, n int) map[int][5][5]int {
	grid := map[int][5][5]int{0: base}
	for i := 0; i < n; i++ {
		newState := make(map[int][5][5]int)
		for z := -n; z <= n; z++ {
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if y == 2 && x == 2 {
						continue
					}
					aliveNeighbors := 0
					pos := image.Point{x, y}
					for _, neighbor := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
						newPos := pos.Add(neighbor)
						if newPos.X >= 5 || newPos.X < 0 || newPos.Y >= 5 || newPos.Y < 0 {
							aliveNeighbors += grid[z-1][neighbor.Y+2][neighbor.X+2]
						} else if newPos.X == 2 && newPos.Y == 2 {
							for j := 0; j < 5; j++ {
								switch neighbor {
								case image.Point{0, -1}:
									aliveNeighbors += grid[z+1][4][j]
								case image.Point{1, 0}:
									aliveNeighbors += grid[z+1][j][0]
								case image.Point{0, 1}:
									aliveNeighbors += grid[z+1][0][j]
								default:
									aliveNeighbors += grid[z+1][j][4]
								}
							}
						} else {
							aliveNeighbors += grid[z][newPos.Y][newPos.X]
						}
					}
					section := newState[z]
					if grid[z][y][x] == BUGS {
						if aliveNeighbors == 1 {
							section[y][x] = BUGS
						} else {
							section[y][x] = SPACE
						}
					} else {
						if aliveNeighbors == 1 || aliveNeighbors == 2 {
							section[y][x] = BUGS
						} else {
							section[y][x] = SPACE
						}
					}
					newState[z] = section
				}
			}
		}
		grid = newState
	}
	return grid
}

func countBugs(grid map[int][5][5]int) int {
	var count int
	for _, section := range grid {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				count += section[y][x]
			}
		}
	}
	return count
}

func main() {
	base := readBase("./input.txt")
	grid := timestep(base, 200)
	fmt.Println(countBugs(grid))
}
