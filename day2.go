package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type opCode int

const (
	add      opCode = 1
	multiply opCode = 2
	halt     opCode = 99
)

func runIntcode(program []int) {
	var op opCode
	var res, pos1, pos2, pos3 int

	for i := 0; ; i += 4 {
		op = opCode(program[i])
		if op == halt {
			return
		}

		pos1, pos2, pos3 = program[i+1], program[i+2], program[i+3]
		switch op {
		case add:
			res = program[pos1] + program[pos2]
		case multiply:
			res = program[pos1] * program[pos2]
		}
		program[pos3] = res
	}
}

func convertToProgram(content []byte) []int {
	input := strings.Split(string(content), ",")
	program := make([]int, len(input), len(input))
	for i, v := range input {
		num, _ := strconv.ParseInt(v, 10, 0)
		program[i] = int(num)
	}
	return program
}

func main() {
	content, _ := ioutil.ReadFile("day2-input.txt")
	program := convertToProgram(content)
	program[1] = 12
	program[2] = 2
	fmt.Println("IntCode Program:", program)

	runIntcode(program)
	fmt.Println("Program after running:", program)
	fmt.Println("Result:", program[0])
}
