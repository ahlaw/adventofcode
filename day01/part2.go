package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func requiredFuel(mass int) int {
	totalFuel := 0
	for {
		mass = mass/3 - 2
		if mass < 0 {
			break
		}
		totalFuel += mass
	}
	return totalFuel
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel += requiredFuel(mass)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fuel)
}
