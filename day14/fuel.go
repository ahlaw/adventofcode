package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Chemical struct {
	Name   string
	Amount int
}

type Reaction struct {
	Product  Chemical
	Reagents []Chemical
}

func oreRequired(reactions map[string]Reaction, target string, targetAmount int, excess map[string]int) int {
	if target == "ORE" {
		return targetAmount
	} else if targetAmount <= excess[target] {
		excess[target] -= targetAmount
		return 0
	}
	targetAmount -= excess[target]
	excess[target] = 0
	var ore int
	outputAmount, inputs := reactions[target].Product.Amount, reactions[target].Reagents
	factor := int(math.Ceil(float64(targetAmount) / float64(outputAmount)))
	for _, input := range inputs {
		ore += oreRequired(reactions, input.Name, input.Amount*factor, excess)
	}
	excess[target] += outputAmount*factor - targetAmount
	return ore
}

func maxFuel(reactions map[string]Reaction, oreCargo int) int {
	low, high, oreAmount := 1, oreCargo, 0
	for low < high {
		mid := (low + high) / 2
		excess := make(map[string]int)
		oreAmount = oreRequired(reactions, "FUEL", mid, excess)
		if oreAmount > oreCargo {
			high = mid - 1
		} else if oreAmount < oreCargo {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)
	reactions := make(map[string]Reaction)
	for sc.Scan() {
		reaction := strings.Split(strings.TrimSpace(sc.Text()), " => ")
		output := strings.Split(reaction[1], " ")
		amount, _ := strconv.Atoi(output[0])
		product := Chemical{Name: output[1], Amount: amount}
		reagents := make([]Chemical, 0)
		for _, s := range strings.Split(reaction[0], ", ") {
			input := strings.Split(s, " ")
			amount, _ := strconv.Atoi(input[0])
			reagents = append(reagents, Chemical{Name: input[1], Amount: amount})
		}
		reactions[output[1]] = Reaction{Product: product, Reagents: reagents}
	}
	excess := make(map[string]int)
	fmt.Print("Part 1: ")
	fmt.Println(oreRequired(reactions, "FUEL", 1, excess))
	oreCargo := 1000000000000
	fmt.Print("Part 2: ")
	fmt.Println(maxFuel(reactions, oreCargo))
}
