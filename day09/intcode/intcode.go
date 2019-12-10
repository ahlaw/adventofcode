package intcode

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
	RELATIVE  ParamMode = 2
)

const (
	ADD                  Opcode = 1
	MULTIPLY             Opcode = 2
	INPUT                Opcode = 3
	OUTPUT               Opcode = 4
	JUMP_IF_TRUE         Opcode = 5
	JUMP_IF_FALSE        Opcode = 6
	LESS_THAN            Opcode = 7
	EQUALS               Opcode = 8
	RELATIVE_BASE_OFFSET Opcode = 9
	HALT                 Opcode = 99
)

func leftpad(str string, num int) string {
	pad := ""
	if num >= 0 {
		pad = strings.Repeat("0", num)
	}
	return pad + str
}

func opcode(code int64) Opcode {
	text := strconv.FormatInt(code, 10)
	padded := leftpad(text, OPCODE_LENGTH-len(text))
	op, _ := strconv.Atoi(padded[len(padded)-2:])
	return Opcode(op)
}

/*
func params(ops []int64, pos int, numParams int, relBase int64) []int64 {
	code := strconv.FormatInt(ops[pos], 10)
	fullLength := OPCODE_LENGTH + numParams
	fullCode := leftpad(code, fullLength-len(code))
	p := make([]int64, numParams)
	for i, mode := range fullCode[:numParams] {
		// Gets int value of ASCII representation of rune
		index := numParams - i - 1
		switch ParamMode(mode - '0') {
		case POSITION:
			p[index] = ops[pos+index+1]
		case IMMEDIATE:
			p[index] = int64(pos + index + 1)
		case RELATIVE:
			p[index] = ops[pos+index+1] + relBase
		}
	}
	return p
}
*/

func RunProgram(ops []int64) error {
	var i int
	var relativeBase int64

	params := func(ops []int64, numParams int) []int64 {
		code := strconv.FormatInt(ops[i], 10)
		fullLength := OPCODE_LENGTH + numParams
		fullCode := leftpad(code, fullLength-len(code))
		p := make([]int64, numParams)
		for codeIndex, mode := range fullCode[:numParams] {
			// Gets int value of ASCII representation of rune
			index := numParams - codeIndex - 1
			switch ParamMode(mode - '0') {
			case POSITION:
				p[index] = ops[i+index+1]
			case IMMEDIATE:
				p[index] = int64(i + index + 1)
			case RELATIVE:
				p[index] = ops[i+index+1] + relativeBase
			}
		}
		return p
	}

	for {
		switch opcode(ops[i]) {
		case ADD:
			p := params(ops, 3)
			ops[p[2]] = ops[p[0]] + ops[p[1]]
			i += 4
		case MULTIPLY:
			p := params(ops, 3)
			ops[p[2]] = ops[p[0]] * ops[p[1]]
			i += 4
		case INPUT:
			fmt.Print("Enter input: ")
			var input int
			fmt.Scanln(&input)
			p := params(ops, 1)
			ops[p[0]] = int64(input)
			i += 2
		case OUTPUT:
			p := params(ops, 1)
			fmt.Println(ops[p[0]])
			i += 2
		case JUMP_IF_TRUE:
			p := params(ops, 2)
			if ops[p[0]] != 0 {
				i = int(ops[p[1]])
			} else {
				i += 3
			}
		case JUMP_IF_FALSE:
			p := params(ops, 2)
			if ops[p[0]] == 0 {
				i = int(ops[p[1]])
			} else {
				i += 3
			}
		case LESS_THAN:
			p := params(ops, 3)
			if ops[p[0]] < ops[p[1]] {
				ops[p[2]] = 1
			} else {
				ops[p[2]] = 0
			}
			i += 4
		case EQUALS:
			p := params(ops, 3)
			if ops[p[0]] == ops[p[1]] {
				ops[p[2]] = 1
			} else {
				ops[p[2]] = 0
			}
			i += 4
		case RELATIVE_BASE_OFFSET:
			p := params(ops, 1)
			relativeBase += ops[p[0]]
			i += 2
		case HALT:
			return nil
		}
	}
	return errors.New("Something went wrong following opcodes")
}

func ReadProgram() []int64 {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(data[:len(data)-1]), ",")
	memory := 3000
	ops := make([]int64, memory)
	for i, num := range input {
		ops[i], err = strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ops
}
