package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"unicode"
)

func isPortal(str string) bool {
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func generatePairs(grid []string, i, j int) []string {
	pair1 := string(grid[i][j-2 : j])
	pair2 := string(grid[i][j+1 : j+3])
	pair3 := string([]byte{grid[i-2][j], grid[i-1][j]})
	pair4 := string([]byte{grid[i+1][j], grid[i+2][j]})
	return []string{pair1, pair2, pair3, pair4}
}

func readGrid(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func readPortals(grid []string) (map[image.Point]image.Point, image.Point, image.Point) {
	portalLoc := make(map[string]image.Point)
	portals := make(map[image.Point]image.Point)
	for i := 2; i < len(grid)-2; i++ {
		for j := 2; j < len(grid[0])-2; j++ {
			if grid[i][j] != '.' {
				continue
			}
			for _, pair := range generatePairs(grid, i, j) {
				if isPortal(pair) {
					currPos := image.Point{i, j}
					if _, ok := portalLoc[pair]; ok {
						portals[currPos] = portalLoc[pair]
						portals[portalLoc[pair]] = currPos
					} else {
						portalLoc[pair] = currPos
					}
					continue
				}
			}
		}
	}
	start := portalLoc["AA"]
	end := portalLoc["ZZ"]
	return portals, start, end
}

func bfs(grid []string, portals map[image.Point]image.Point, start, end image.Point) int {
	var queue = []image.Point{start}
	dist := map[image.Point]int{start: 0}
	for len(queue) > 0 {
		currPos := queue[0]
		queue = queue[1:]
		if currPos == end {
			return dist[currPos]
		}
		for _, move := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, 0}} {
			newPos := currPos.Add(move)
			var ok bool
			if move.Eq(image.Point{0, 0}) {
				if newPos, ok = portals[currPos]; !ok {
					continue
				}
			}
			if _, ok := dist[newPos]; ok {
				continue
			}
			if grid[newPos.X][newPos.Y] == '.' {
				dist[newPos] = dist[currPos] + 1
				queue = append(queue, newPos)
			}
		}
	}
	return -1
}

func main() {
	grid := readGrid("./input.txt")
	portals, start, end := readPortals(grid)
	fmt.Println(bfs(grid, portals, start, end))
}
