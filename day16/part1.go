package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func fft(signal []int, phases int) []int {
	for p := 0; p < phases; p++ {
		output := make([]int, len(signal))
		for i := range signal {
			sum := 0
			for j, value := range signal {
				sum += value * []int{0, 1, 0, -1}[(j+1)/(i+1)%4]
			}
			if sum < 0 {
				sum = -sum
			}
			output[i] = sum % 10
		}
		signal = output
	}
	return signal
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = bytes.TrimSpace(input)
	signal := make([]int, len(input))
	for i, c := range input {
		signal[i] = int(c - '0')
	}
	fmt.Print("Part 1: ")
	for _, i := range fft(signal, 100)[:8] {
		fmt.Print(i)
	}
	fmt.Println()
}
