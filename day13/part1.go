package main

import (
	"fmt"
	"github.com/ahlaw/adventofcode/day13/intcode"
)

func main() {
	ops := intcode.ReadProgram()
	in := make(chan int, 1)
	out := make(chan int)
	go intcode.RunProgram(ops, in, out)

	count := 0
	for {
		_, ok := <-out
		if !ok {
			break
		}
		<-out
		if <-out == 2 {
			count++
		}
	}
	fmt.Println(count)
}
