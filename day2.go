package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	noun         = 1
	verb         = 2
	searchOutput = 19690720
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
	stringContent := strings.TrimSpace(string(content))
	return strings.Split(stringContent, ",")
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
	initialMemory := initMemory(intcode)

	memory := make([]int, len(intcode), len(intcode))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(memory, initialMemory)
			memory[1] = noun
			memory[2] = verb
			fmt.Println("Initial memory:", memory)
			runIntcode(memory)
			fmt.Println("Memory after running Intcode:", memory)
			output := memory[0]
			if output == searchOutput {
				fmt.Println("Noun, verb:", noun, verb)
				fmt.Println("Result:", 100*noun+verb)
				return
			}
		}
	}
}
