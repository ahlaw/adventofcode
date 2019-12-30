package main

import (
	"bufio"
	"fmt"
	"github.com/ahlaw/adventofcode/day21/intcode"
	"io"
	"math"
	"os"
)

type springdroid struct {
	program []int64
}

func readLine(out <-chan int) string {
	var line []rune
	for c := range out {
		if c == '\n' {
			break
		}
		line = append(line, rune(c))
	}
	return string(line)
}

func writeLine(in chan<- int, line string) {
	for _, c := range line {
		in <- int(c)
	}
	in <- int('\n')
}

func (d springdroid) run(r io.Reader, prompt bool) {
	in := make(chan int)
	out := make(chan int)
	go intcode.RunProgram(d.program, in, out)
	line := readLine(out)
	if prompt {
		fmt.Println(line)
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		writeLine(in, line)
		if line == "WALK" || line == "RUN" {
			break
		}
	}
	for c := range out {
		if c > math.MaxInt8 {
			fmt.Println(c)
		} else {
			fmt.Printf("%c", c)
		}
	}
}

func main() {
	ops := intcode.ReadProgram()
	droid := springdroid{ops}
	if len(os.Args) == 2 {
		script, _ := os.Open(os.Args[1])
		defer script.Close()
		droid.run(script, false)
	} else {
		droid.run(os.Stdin, true)
	}
}
