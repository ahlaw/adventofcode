package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func mapOrbit(input []string) map[string]string {
	orbits := make(map[string]string)
	for _, line := range input {
		bodies := strings.Split(line, ")")
		orbits[bodies[1]] = bodies[0]
	}
	return orbits
}

func countOrbits(orbits map[string]string, body string) (count int) {
	for body != "COM" {
		body = orbits[body]
		count++
	}
	return
}

func main() {
	input := readInput()
	orbits := mapOrbit(input)
	var count int
	for body, _ := range orbits {
		count += countOrbits(orbits, body)
	}
	fmt.Println(count)
}
