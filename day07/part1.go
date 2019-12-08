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
	state := make([]int, len(ops))
	ch := make(chan int)

	highest := 0
	for _, p := range permutation([]int{0, 1, 2, 3, 4}) {
		output := 0
		for _, phase := range p {
			copy(state, ops)
			go intcode.Run(state, ch)
			ch <- phase
			ch <- output
			output = <-ch
		}
		if output > highest {
			highest = output
		}
	}
	fmt.Println(highest)
}
