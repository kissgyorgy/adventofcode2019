package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type opCode int

const (
	add      opCode = 1
	multiply opCode = 2
	input    opCode = 3
	output   opCode = 4
	halt     opCode = 99
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
	return val
}

func runIntcode(memory []int, inputVal int) {
	var op opCode
	var param1, param2, respos int

	for addr := 0; ; {
		fmt.Println("Instruction:", memory[addr])
		op = opCode(memory[addr] % 100)
		if op == halt {
			return
		}

		param1 = getParam(memory, addr, 1)

		if op == add || op == multiply {
			param2 = getParam(memory, addr, 2)
			// Parameters that an instruction writes to will never be in immediate mode.
			respos = memory[addr+3]
		}

		switch op {
		case add:
			fmt.Printf("ADD: %v+%v => %v\n", param1, param2, respos)
			memory[respos] = param1 + param2
			addr += 4
		case multiply:
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
		}
	}
}

func main() {
	intcode := loadProgram("day5-input.txt")
	memory := initMemory(intcode)
	fmt.Println("Memory:", memory)
	inputVal := 1
	runIntcode(memory, inputVal)
}
