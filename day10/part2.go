package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Coord struct {
	X, Y int
}

type Vector struct {
	X, Y int
}

func (a Coord) distance(b Coord) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return 1
	} else if a < 0 {
		return -a
	}
	return a
}

func locateAsteroids(input []string) []Coord {
	asteroids := make([]Coord, 0)
	rows, cols := len(input), len(input[0])
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if input[y][x] == '#' {
				asteroids = append(asteroids, Coord{x, y})
			}
		}
	}
	return asteroids
}

func calculateSlopes(asteroids []Coord) map[Coord][]Vector {
	slopeMap := make(map[Coord][]Vector)
	for _, base := range asteroids {
		slopes := make([]Vector, 0)
		for _, other := range asteroids {
			if base != other {
				slope := Vector{base.X - other.X, base.Y - other.Y}
				slopes = append(slopes, slope)
			}

		}
		slopeMap[base] = slopes
	}
	return slopeMap
}

func simplifySlopes(slopeMap map[Coord][]Vector) map[Coord][]Vector {
	newMap := make(map[Coord][]Vector)
	for asteroid := range slopeMap {
		tmp := make([]Vector, len(slopeMap[asteroid]))
		copy(tmp, slopeMap[asteroid])
		for i, slope := range slopeMap[asteroid] {
			div := gcd(slope.X, slope.Y)
			newSlope := Vector{slope.X / div, slope.Y / div}
			tmp[i] = newSlope
		}
		newMap[asteroid] = tmp
	}
	return newMap
}

func countIndependent(slopeMap map[Coord][]Vector) map[Coord]int {
	asteroidCounts := make(map[Coord]int)
	for asteroid, slopes := range slopeMap {
		count := 0
		counter := make(map[Vector]int)
		for _, slope := range slopes {
			if counter[slope] == 0 {
				count++
			}
			counter[slope]++
		}
		asteroidCounts[asteroid] = count
	}
	return asteroidCounts
}

func bestAsteroid(counts map[Coord]int) (Coord, int) {
	var best Coord
	var max int
	for asteroid, count := range counts {
		if count > max {
			best = asteroid
			max = count
		}
	}
	return best, max
}

func mapAngles(asteroids []Coord, base Coord) map[float64][]Coord {
	angleMap := make(map[float64][]Coord)
	for _, other := range asteroids {
		if base != other {
			relative := Coord{other.X - base.X, other.Y - base.Y}
			angle := -(math.Atan2(float64(relative.X), float64(relative.Y)) + math.Pi) + 2*math.Pi
			angleMap[angle] = append(angleMap[angle], other)
		}

	}
	// Sort asteroids by euclidean distance
	for _, asteroids := range angleMap {
		sort.Slice(asteroids, func(i, j int) bool {
			return asteroids[i].distance(base) < asteroids[j].distance(base)
		})
	}
	return angleMap
}

func destroyAsteroidNum(angleMap map[float64][]Coord, targetNum int) Coord {
	// Get sorted angles
	var angles []float64
	for angle := range angleMap {
		angles = append(angles, angle)
	}
	sort.Float64s(angles)
	i := 0
	for i <= targetNum {
		for _, angle := range angles {
			if len(angleMap[angle]) > 0 {
				i++
				if i == targetNum {
					return angleMap[angle][0]
				}
				angleMap[angle] = angleMap[angle][1:]
			}
		}
	}
	return Coord{0, 0}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	asteroids := locateAsteroids(input)
	slopeMap := calculateSlopes(asteroids)
	reducedSlopes := simplifySlopes(slopeMap)
	visibleCounts := countIndependent(reducedSlopes)
	best, _ := bestAsteroid(visibleCounts)
	fmt.Println("Best asteroid:", best)
	angleMap := mapAngles(asteroids, best)
	target := destroyAsteroidNum(angleMap, 200)
	fmt.Println("200th asteroid to be destroyed:", target)
	fmt.Println(target.X*100 + target.Y)
}
