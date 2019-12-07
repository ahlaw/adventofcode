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

func orbitPath(orbits map[string]string, body string) []string {
	path := make([]string, 0)
	for body != "COM" {
		body = orbits[body]
		path = append(path, body)
	}
	return path
}

func symmetricDifference(arr1, arr2 []string) (diff []string) {
	m := make(map[string]bool)
	n := make(map[string]bool)

	for _, item := range arr1 {
		m[item] = true
	}

	for _, item := range arr2 {
		n[item] = true
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	for _, item := range arr1 {
		if _, ok := n[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func main() {
	input := readInput()
	orbits := mapOrbit(input)
	p1 := orbitPath(orbits, "YOU")
	p2 := orbitPath(orbits, "SAN")
	fmt.Println(len(symmetricDifference(p1, p2)))
}
