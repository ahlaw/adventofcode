package main

import (
	"bufio"
	"fmt"
	"github.com/ahlaw/adventofcode/day25/intcode"
	"io"
	"os"
)

const (
	prompt = "Command?"
)

type droid struct {
	program []int64
}

func readLine(out <-chan int) (string, bool) {
	var line []rune
	var ok bool
	var c int
	for c, ok = <-out; ok && c != '\n'; c = <-out {
		line = append(line, rune(c))
	}
	return string(line), ok
}

func writeLine(in chan<- int, line string) {
	for _, c := range line {
		in <- int(c)
	}
	in <- int('\n')
}

func getCommand(sc *bufio.Scanner) (string, bool) {
	if sc.Scan() {
		return sc.Text(), true
	}
	return "", false
}

func (d droid) run(r io.Reader) {
	in := make(chan int)
	out := make(chan int)
	go intcode.RunProgram(d.program, in, out)
	sc := bufio.NewScanner(r)
	for {
		line, ok := readLine(out)
		if !ok {
			fmt.Println("Machine halted. Exiting...")
			return
		}
		fmt.Println(line)
		if line != prompt {
			continue
		}
		cmd, ok := getCommand(sc)
		if !ok {
			fmt.Println("No more commands. Exiting...")
			return
		}
		writeLine(in, cmd)
	}
}

func main() {
	ops := intcode.ReadProgram()
	d := droid{ops}
	d.run(os.Stdin)
}
