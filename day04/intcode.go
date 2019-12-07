package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const OPCODE_LENGTH = 2

type ParamMode int
type Opcode int

const (
	POSITION  ParamMode = 0
	IMMEDIATE ParamMode = 1
)

const (
	ADD           Opcode = 1
	MULTIPLY      Opcode = 2
	INPUT         Opcode = 3
	OUTPUT        Opcode = 4
	JUMP_IF_TRUE  Opcode = 5
	JUMP_IF_FALSE Opcode = 6
	LESS_THAN     Opcode = 7
	EQUALS        Opcode = 8
	HALT          Opcode = 99
)

func leftpad(str string, num int) string {
	pad := ""
	if num >= 0 {
		pad = strings.Repeat("0", num)
	}
	return pad + str
}

func opcode(code int) Opcode {
	text := strconv.Itoa(code)
	padded := leftpad(text, OPCODE_LENGTH-len(text))
	op, _ := strconv.Atoi(padded[len(padded)-2:])
	return Opcode(op)
}

func params(ops []int, pos int, numParams int) []int {
	code := strconv.Itoa(ops[pos])
	fullLength := OPCODE_LENGTH + numParams
	fullCode := leftpad(code, fullLength-len(code))
	p := make([]int, numParams)
	for i, mode := range fullCode[:numParams] {
		// Gets int value of ASCII representation of rune
		index := numParams - i - 1
		switch ParamMode(mode - '0') {
		case POSITION:
			p[index] = ops[pos+index+1]
		case IMMEDIATE:
			p[index] = pos + index + 1
		}
	}
	return p
}

func runProgram(ops []int) error {
	i := 0
	for {
		switch opcode(ops[i]) {
		case ADD:
			p := params(ops, i, 3)
			ops[p[2]] = ops[p[0]] + ops[p[1]]
			i += 4
		case MULTIPLY:
			p := params(ops, i, 3)
			ops[p[2]] = ops[p[0]] * ops[p[1]]
			i += 4
		case INPUT:
			fmt.Print("Enter system ID: ")
			var input int
			fmt.Scanln(&input)
			p := params(ops, i, 1)
			ops[p[0]] = int(input)
			i += 2
		case OUTPUT:
			p := params(ops, i, 1)
			fmt.Println(ops[p[0]])
			i += 2
		case JUMP_IF_TRUE:
			p := params(ops, i, 2)
			if ops[p[0]] != 0 {
				i = ops[p[1]]
			} else {
				i += 3
			}
		case JUMP_IF_FALSE:
			p := params(ops, i, 2)
			if ops[p[0]] == 0 {
				i = ops[p[1]]
			} else {
				i += 3
			}
		case LESS_THAN:
			p := params(ops, i, 3)
			if ops[p[0]] < ops[p[1]] {
				ops[p[2]] = 1
			} else {
				ops[p[2]] = 0
			}
			i += 4
		case EQUALS:
			p := params(ops, i, 3)
			if ops[p[0]] == ops[p[1]] {
				ops[p[2]] = 1
			} else {
				ops[p[2]] = 0
			}
			i += 4
		case HALT:
			return nil
		}
	}
	return errors.New("Something went wrong following opcodes")
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
	err = runProgram(ops)
	if err != nil {
		log.Fatal(err)
	}
}
