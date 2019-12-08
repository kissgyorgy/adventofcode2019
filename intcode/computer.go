package intcode

import (
	"fmt"
	"log"
	"os"
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

func Run(name string, memory []int, inputs, outputs chan int) {
	var op opCode
	var param1, param2, respos int
	prefix := fmt.Sprintf("[%s] ", name)
	l := log.New(os.Stdout, prefix, 0)

	for addr := 0; ; {
		l.Printf("Instruction: %v\n", memory[addr])
		l.Printf("Addr: %v\n", addr)
		op = opCode(memory[addr] % 100)
		if op == halt {
			close(outputs)
			return
		}

		param1 = getParam(l, memory, addr, 1)

		switch op {
		case add:
			param2 = getParam(l, memory, addr, 2)
			// Parameters that an instruction writes to will never be in immediate mode.
			respos = memory[addr+3]
			l.Printf("ADD: %d+%d => %d\n", param1, param2, respos)
			memory[respos] = param1 + param2
			addr += 4

		case multiply:
			param2 = getParam(l, memory, addr, 2)
			// Parameters that an instruction writes to will never be in immediate mode.
			respos = memory[addr+3]
			l.Printf("MUL: %d*%d => %d\n", param1, param2, respos)
			memory[respos] = param1 * param2
			addr += 4

		case input:
			respos = memory[addr+1]
			l.Printf("Waiting for INPUT: ? <= %d\n", respos)
			in := <-inputs
			l.Printf("Got INPUT: %d => %d\n", in, respos)
			memory[respos] = in
			addr += 2

		case output:
			l.Printf("OUTPUT: %d \n", param1)
			outputs <- param1
			addr += 2

		case jumpIfTrue:
			if param1 != 0 {
				param2 = getParam(l, memory, addr, 2)
				l.Printf("JUMP: => %d\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case jumpIfFalse:
			if param1 == 0 {
				param2 = getParam(l, memory, addr, 2)
				l.Printf("JUMP: => %d\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case lessThan:
			param2 = getParam(l, memory, addr, 2)
			l.Println("LESSTHAN:", param1, param2)
			respos = memory[addr+3]
			if param1 < param2 {
				memory[respos] = 1
			} else {
				memory[respos] = 0
			}
			addr += 4

		case equals:
			param2 = getParam(l, memory, addr, 2)
			l.Println("EQUALS:", param1, param2)
			respos = memory[addr+3]
			if param1 == param2 {
				memory[respos] = 1
			} else {
				memory[respos] = 0
			}
			addr += 4

		default:
			l.Println("Invalid opcode:", op)
			os.Exit(1)
		}
	}
}
