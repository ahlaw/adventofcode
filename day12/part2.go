package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	X = 0
	Y = 1
	Z = 2
)

type Vector [3]int

type Moon struct {
	Pos, Vel Vector
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(ints ...int64) int64 {
	if len(ints) > 2 {
		return gcd(ints[0], gcd(ints[1:]...))
	}
	a, b := ints[0], ints[1]
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(ints ...int64) int64 {
	if len(ints) > 2 {
		return lcm(ints[0], lcm(ints[1:]...))
	}
	a, b := ints[0], ints[1]
	if a == 0 && b == 0 {
		return 0
	}
	return abs(a*b) / gcd(a, b)
}

func (m1 *Moon) ApplyGravity(m2 Moon, ax int) {
	if m1.Pos[ax] < m2.Pos[ax] {
		m1.Vel[ax]++
	} else if m1.Pos[ax] > m2.Pos[ax] {
		m1.Vel[ax]--
	}
}

func (m *Moon) ApplyVelocity(ax int) {
	m.Pos[ax] += m.Vel[ax]
}

func step(moons []Moon, ax int) {
	for i := range moons {
		for j := range moons {
			if i != j {
				moons[i].ApplyGravity(moons[j], ax)
			}
		}
	}
	for i := range moons {
		moons[i].ApplyVelocity(ax)
	}
}

func completePeriod(moons []Moon, moonsInit []Moon, ax int) bool {
	for i := range moons {
		if moons[i].Pos[ax] != moonsInit[i].Pos[ax] || moons[i].Vel[ax] != moonsInit[i].Vel[ax] {
			return false
		}
	}
	return true
}

func periodLength(moonsInit []Moon, ax int) int64 {
	moons := make([]Moon, len(moonsInit))
	copy(moons, moonsInit)
	var count int64
	for {
		step(moons, ax)
		count++
		if completePeriod(moons, moonsInit, ax) {
			break
		}
	}
	return count
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)
	var moons []Moon
	for sc.Scan() {
		moon := Moon{}
		fmt.Sscanf(sc.Text(), "<x=%d, y=%d, z=%d>", &moon.Pos[X], &moon.Pos[Y], &moon.Pos[Z])
		moons = append(moons, moon)
	}
	periods := make([]int64, 3)
	for ax := range []int{X, Y, Z} {
		periods[ax] = periodLength(moons, ax)
	}
	fmt.Println(lcm(periods...))
}
