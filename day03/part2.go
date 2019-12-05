package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, steps int
}

func points(path []string) []Point {
	curr := Point{}
	steps := 0
	p := make([]Point, 0)
	for _, step := range path {
		dir := step[0]
		dist, err := strconv.Atoi(step[1:])
		if err != nil {
			log.Fatal(err)
		}
		stepPath := make([]Point, dist)
		for i := 0; i < dist; i++ {
			steps++
			switch dir {
			case 'U':
				stepPath[i] = Point{curr.X, curr.Y + 1, steps}
				curr.Y++
			case 'D':
				stepPath[i] = Point{curr.X, curr.Y - 1, steps}
				curr.Y--
			case 'L':
				stepPath[i] = Point{curr.X - 1, curr.Y, steps}
				curr.X--
			case 'R':
				stepPath[i] = Point{curr.X + 1, curr.Y, steps}
				curr.X++
			}
		}
		p = append(p, stepPath...)
	}
	return p
}

func bestIntersect(set1, set2 []Point) int {
	var result int
	intersects := false
	for _, p1 := range set1 {
		for _, p2 := range set2 {
			if p1.X == p2.X && p1.Y == p2.Y {
				steps := p1.steps + p2.steps
				if intersects == false || steps < result {
					result = steps
					intersects = true
				}
			}
		}
	}
	return result
}

func readInput() (path1, path2 []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	paths := make([][]string, 0)
	for scanner.Scan() {
		paths = append(paths, strings.Split(string(scanner.Text()), ","))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	path1, path2 = paths[0], paths[1]
	return
}

func main() {
	path1, path2 := readInput()
	set1 := points(path1)
	set2 := points(path2)
	fmt.Println(bestIntersect(set1, set2))
}
