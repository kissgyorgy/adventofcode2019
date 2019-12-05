package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	IntcodeFile = "day5-input.txt"
	inputVal    = 5
)

type opCode int

const (
	add         opCode = 1
	multiply    opCode = 2
	input       opCode = 3
	output      opCode = 4
	jumpIfTrue  opCode = 5
	jumpIfFalse opCode = 6
	lessThan    opCode = 7
	equals      opCode = 8
	halt        opCode = 99
)

type paramMode int

const (
	positionMode  paramMode = 0
	immediateMode paramMode = 1
)

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

// get nth digit of number counted from rigth to left
func getNthDigitFromRight(num, ind int) int {
	return num / int(math.Pow10(ind)) % 10
}

func getParam(memory []int, opAddr, nth int) int {
	var val int
	mode := getNthDigitFromRight(memory[opAddr], nth+1)
	switch paramMode(mode) {
	case positionMode:
		pos := memory[opAddr+nth]
		val = memory[pos]
	case immediateMode:
		val = memory[opAddr+nth]
	default:
		fmt.Println("Invalid parameter mode:", mode)
		os.Exit(1)
	}
	fmt.Printf("param%d: %v (addr: %d)\n", nth, val, opAddr+nth)
	return val
}

func runIntcode(memory []int, inputVal int) {
	var op opCode
	var param1, param2, respos int

	for addr := 0; ; {
		fmt.Println("----")
		fmt.Printf("Instruction: %v\n", memory[addr])
		fmt.Println("Addr", addr)
		op = opCode(memory[addr] % 100)
		if op == halt {
			return
		}

		param1 = getParam(memory, addr, 1)

		switch op {
		case add:
			param2 = getParam(memory, addr, 2)
			// Parameters that an instruction writes to will never be in immediate mode.
			respos = memory[addr+3]
			fmt.Printf("ADD: %v+%v => %v\n", param1, param2, respos)
			memory[respos] = param1 + param2
			addr += 4

		case multiply:
			param2 = getParam(memory, addr, 2)
			// Parameters that an instruction writes to will never be in immediate mode.
			respos = memory[addr+3]
			fmt.Printf("MUL: %v*%v => %v\n", param1, param2, respos)
			memory[respos] = param1 * param2
			addr += 4

		case input:
			respos = memory[addr+1]
			fmt.Printf("INPUT: %v => %v\n", inputVal, respos)
			memory[respos] = inputVal
			addr += 2

		case output:
			fmt.Printf("OUT: %d \n", param1)
			addr += 2

		case jumpIfTrue:
			if param1 != 0 {
				param2 = getParam(memory, addr, 2)
				fmt.Printf("JUMP: => %v\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case jumpIfFalse:
			if param1 == 0 {
				param2 = getParam(memory, addr, 2)
				fmt.Printf("JUMP: => %v\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case lessThan:
			param2 = getParam(memory, addr, 2)
			fmt.Println("LESSTHAN:", param1, param2)
			respos = memory[addr+3]
			if param1 < param2 {
				memory[respos] = 1
			} else {
				memory[respos] = 0
			}
			addr += 4

		case equals:
			param2 = getParam(memory, addr, 2)
			fmt.Println("EQUALS:", param1, param2)
			respos = memory[addr+3]
			if param1 == param2 {
				memory[respos] = 1
			} else {
				memory[respos] = 0
			}
			addr += 4

		default:
			fmt.Println("Invalid opcode:", op)
			os.Exit(1)
		}
	}
}

func main() {
	intcode := loadProgram(IntcodeFile)
	memory := initMemory(intcode)
	fmt.Println("Memory:", memory)
	runIntcode(memory, inputVal)
}
