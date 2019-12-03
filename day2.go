package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	noun = 1
	verb = 2
)

type opCode int

const (
	add      opCode = 1
	multiply opCode = 2
	halt     opCode = 99
)

func runIntcode(program []int) {
	var op opCode
	var param1, param2, param3 int
	var value1, value2, res int

	for addr := 0; ; addr += 4 {
		instr := program[addr : addr+4]

		op = opCode(instr[0])
		if op == halt {
			return
		}

		param1, param2, param3 = instr[1], instr[2], instr[3]
		value1, value2 = program[param1], program[param2]
		switch op {
		case add:
			res = value1 + value2
		case multiply:
			res = value1 * value2
		}
		program[param3] = res
	}
}

func loadProgram(filename string) []string {
	content, _ := ioutil.ReadFile(filename)
	return strings.Split(string(content), ",")
}

func initMemory(input []string) []int {
	memory := make([]int, len(input), len(input))
	for i, v := range input {
		num, _ := strconv.ParseInt(v, 10, 0)
		memory[i] = int(num)
	}
	return memory
}

func main() {
	intcode := loadProgram("day2-input.txt")
	memory := initMemory(intcode)

	noun, verb := 12, 2
	memory[1] = noun
	memory[2] = verb
	fmt.Println("Initial memory:", memory)

	runIntcode(memory)
	fmt.Println("Memory after running Intcode:", memory)

	output := memory[0]
	fmt.Println("Result:", output)
}
