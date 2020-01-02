package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day23/intcode"
)

type Packet struct {
	X, Y int
}

type Nat struct {
	old, curr Packet
}

func startNetwork(ops []int64) {
	in := make([]chan int, 50)
	out := make([]chan int, 50)
	for i := 0; i < 50; i++ {
		in[i], out[i] = make(chan int), make(chan int)
		prog := make([]int64, len(ops))
		copy(prog, ops)
		go intcode.RunProgram(prog, in[i], out[i])
		in[i] <- i
		in[i] <- -1
	}
	var nat Nat
	var part1 = true
	idle := 0
	for i := 0; ; i = (i + 1) % 50 {
		select {
		case addr := <-out[i]:
			if addr == 255 {
				nat.curr = Packet{<-out[i], <-out[i]}
				if part1 {
					fmt.Print("Part1: ")
					fmt.Println(nat.curr.Y)
					part1 = false
				}
			} else {
				in[addr] <- <-out[i]
				in[addr] <- <-out[i]
			}
			idle = 0
		case in[i] <- -1:
			idle++
		}

		if idle >= 50 {
			if nat.curr.Y == nat.old.Y {
				fmt.Print("Part2: ")
				fmt.Println(nat.curr.Y)
				break
			}
			in[0] <- nat.curr.X
			in[0] <- nat.curr.Y
			nat.old = nat.curr
			idle = 0
		}
	}
}

func main() {
	ops := intcode.ReadProgram()
	startNetwork(ops)
}
