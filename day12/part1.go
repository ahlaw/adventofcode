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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m1 *Moon) ApplyGravity(m2 Moon) {
	for ax := range []int{X, Y, Z} {
		if m1.Pos[ax] < m2.Pos[ax] {
			m1.Vel[ax]++
		} else if m1.Pos[ax] > m2.Pos[ax] {
			m1.Vel[ax]--
		}
	}
}

func (m *Moon) ApplyVelocity() {
	for ax := range []int{X, Y, Z} {
		m.Pos[ax] += m.Vel[ax]
	}
}

func (m *Moon) TotalEnergy() int {
	var pe, ke int
	for ax := range []int{X, Y, Z} {
		pe += abs(m.Pos[ax])
		ke += abs(m.Vel[ax])
	}
	return pe * ke
}

func step(moons []Moon) {
	for i := range moons {
		for j := range moons {
			if i != j {
				moons[i].ApplyGravity(moons[j])
			}
		}
	}
	for i := range moons {
		moons[i].ApplyVelocity()
	}
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
	target := 1000
	for i := 0; i < target; i++ {
		step(moons)
	}
	var sysEnergy int
	for i := range moons {
		sysEnergy += moons[i].TotalEnergy()
	}
	fmt.Println(sysEnergy)
}
