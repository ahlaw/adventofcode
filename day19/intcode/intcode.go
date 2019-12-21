package intcode

import (
	"errors"
	"fmt"
	"io/ioutil"
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

func RunProgram(ops []int64, in chan int, out chan int) error {
	var i int
	var relativeBase int64

	params := func(numParams int) []int64 {
		instruction := fmt.Sprintf("%05d", ops[i])
		p := make([]int64, numParams)
		for codeIndex, mode := range instruction[3-numParams : 3] {
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
		instruction := fmt.Sprintf("%05d", ops[i])
		op, _ := strconv.Atoi(instruction[3:])
		switch Opcode(op) {
		case ADD:
			p := params(3)
			ops[p[2]] = ops[p[0]] + ops[p[1]]
			i += 4
		case MULTIPLY:
			p := params(3)
			ops[p[2]] = ops[p[0]] * ops[p[1]]
			i += 4
		case INPUT:
			p := params(1)
			ops[p[0]] = int64(<-in)
			i += 2
		case OUTPUT:
			p := params(1)
			out <- int(ops[p[0]])
			i += 2
		case JUMP_IF_TRUE:
			p := params(2)
			if ops[p[0]] != 0 {
				i = int(ops[p[1]])
			} else {
				i += 3
			}
		case JUMP_IF_FALSE:
			p := params(2)
			if ops[p[0]] == 0 {
				i = int(ops[p[1]])
			} else {
				i += 3
			}
		case LESS_THAN:
			p := params(3)
			if ops[p[0]] < ops[p[1]] {
				ops[p[2]] = 1
			} else {
				ops[p[2]] = 0
			}
			i += 4
		case EQUALS:
			p := params(3)
			if ops[p[0]] == ops[p[1]] {
				ops[p[2]] = 1
			} else {
				ops[p[2]] = 0
			}
			i += 4
		case RELATIVE_BASE_OFFSET:
			p := params(1)
			relativeBase += ops[p[0]]
			i += 2
		case HALT:
			close(out)
			return nil
		}
	}
	return errors.New("Something went wrong following opcodes")
}

func ReadProgram() []int64 {
	data, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(data[:len(data)-1]), ",")
	memory := 3000
	ops := make([]int64, memory)
	for i, num := range input {
		ops[i], _ = strconv.ParseInt(num, 10, 64)
	}
	return ops
}
