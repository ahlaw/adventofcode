package main

import (
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
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	input := strings.Split(string(data[:len(data)-1]), ",")
	ops := make([]int, 0, len(input))
	for _, num := range input {
		code, err := strconv.Atoi(num)
		check(err)
		ops = append(ops, code)
	}
	for i := 0; i < len(ops); i += 4 {
		switch Opcode(ops[i]) {
		case ADD:
			ops[ops[i+3]] = ops[ops[i+1]] + ops[ops[i+2]]
		case MULTIPLY:
			ops[ops[i+3]] = ops[ops[i+1]] * ops[ops[i+2]]
		default:
			break
		}
	}
	fmt.Println(ops)
}
