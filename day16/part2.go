package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func fft(signal []int, phases int) []int {
	for p := 0; p < phases; p++ {
		sum := 0
		for i := len(signal) - 1; i >= 0; i-- {
			sum += signal[i]
			signal[i] = sum % 10
		}
	}
	return signal
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	offset, _ := strconv.Atoi(string(input)[:7])
	input = bytes.Repeat(bytes.TrimSpace(input), 10000)
	signal := make([]int, len(input)-offset)
	for i, c := range input[offset:] {
		signal[i] = int(c - '0')
	}
	fmt.Print("Part 2: ")
	for _, i := range fft(signal, 100)[:8] {
		fmt.Print(i)
	}
	fmt.Println()
}
