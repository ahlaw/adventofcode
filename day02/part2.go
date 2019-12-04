package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Opcode int

const (
	ADD      Opcode = 1
	MULTIPLY Opcode = 2
	HALT     Opcode = 99
)

func check(ops []int, noun int, verb int) (int, error) {
	ops[1] = noun
	ops[2] = verb
	for i := 0; i < len(ops); i += 4 {
		switch Opcode(ops[i]) {
		case ADD:
			ops[ops[i+3]] = ops[ops[i+1]] + ops[ops[i+2]]
		case MULTIPLY:
			ops[ops[i+3]] = ops[ops[i+1]] * ops[ops[i+2]]
		case HALT:
			return ops[0], nil
		}
	}
	return 0, errors.New("Something went wrong following opcodes")
}

func search(ops []int) (int, int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			tmp := make([]int, len(ops))
			copy(tmp, ops)
			output, err := check(tmp, noun, verb)
			if err != nil {
				log.Fatal(err)
			}
			if output == 19690720 {
				return noun, verb, nil
			}
		}
	}
	return 0, 0, errors.New("Could not find value")
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(data[:len(data)-1]), ",")
	ops := make([]int, 0, len(input))
	for _, num := range input {
		code, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		ops = append(ops, code)
	}
	noun, verb, err := search(ops)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Noun = %v, Verb = %v\n", noun, verb)
	product := 100*noun + verb
	fmt.Printf("Product = %v\n", product)
}
