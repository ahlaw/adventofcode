package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"unicode"
)

type node struct {
	pos   image.Point
	depth int
}

func isPortal(str string) bool {
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func isInner(grid []string, coord image.Point) bool {
	return (3 <= coord.X && coord.X < len(grid)-3) && (3 <= coord.Y && coord.Y < len(grid[0])-3)
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

func readPortals(grid []string) (map[image.Point]image.Point, node, node) {
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
	start := node{portalLoc["AA"], 0}
	end := node{portalLoc["ZZ"], 0}
	return portals, start, end
}

func bfs(grid []string, portals map[image.Point]image.Point, start, end node) int {
	var queue = []node{start}
	dist := map[node]int{start: 0}
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]
		// max level assumption
		if currNode.depth > 100 {
			continue
		}
		if currNode == end {
			return dist[currNode]
		}
		for _, move := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, 0}} {
			newNode := node{currNode.pos.Add(move), currNode.depth}
			if move.Eq(image.Point{0, 0}) {
				if _, ok := portals[currNode.pos]; !ok {
					continue
				}
				newDepth := currNode.depth
				if isInner(grid, currNode.pos) {
					newDepth++
				} else if newDepth > 0 {
					newDepth--
				} else {
					continue
				}
				newNode = node{portals[currNode.pos], newDepth}
			}
			if _, ok := dist[newNode]; ok {
				continue
			}
			if grid[newNode.pos.X][newNode.pos.Y] == '.' {
				dist[newNode] = dist[currNode] + 1
				queue = append(queue, newNode)
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
