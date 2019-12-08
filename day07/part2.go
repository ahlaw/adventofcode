package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day07/intcode"
)

func permutation(arr []int) [][]int {
	if len(arr) == 1 {
		return [][]int{arr}
	}
	out := [][]int{}
	for i, num := range arr {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		for _, p := range permutation(append(tmp[:i], tmp[i+1:]...)) {
			out = append(out, append([]int{num}, p...))
		}
	}
	return out
}

func main() {
	ops := intcode.ReadProgram()
	phaseSettings := []int{5, 6, 7, 8, 9}
	chans := make([]chan int, len(phaseSettings))
	highest := 0
	for _, p := range permutation(phaseSettings) {
		for i, phase := range p {
			state := make([]int, len(ops))
			copy(state, ops)
			chans[i] = make(chan int)
			go intcode.Run(state, chans[i])
			chans[i] <- phase
		}
		output := 0
		done := false
		for !done {
			for i := range p {
				select {
				case chans[i] <- output:
					output = <-chans[i]
				default:
					done = true
				}
			}
		}
		if output > highest {
			highest = output
		}
	}
	fmt.Println(highest)
}
