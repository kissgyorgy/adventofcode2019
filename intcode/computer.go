package intcode

import (
	"fmt"
	"log"
	"os"
)

type computer struct {
	memory       []int
	relativeBase int
	inputs       chan int
	outputs      chan int
	l            *log.Logger
}

type opCode int

const (
	add                opCode = 1
	multiply           opCode = 2
	input              opCode = 3
	output             opCode = 4
	jumpIfTrue         opCode = 5
	jumpIfFalse        opCode = 6
	lessThan           opCode = 7
	equals             opCode = 8
	relativeBaseOffset opCode = 9
	halt               opCode = 99
)

func new(name string, program []int, inputs, outputs chan int) *computer {
	logPrefix := fmt.Sprintf("[%s] ", name)
	mem := make([]int, len(program))
	copy(mem, program)
	return &computer{
		memory:       mem,
		relativeBase: 0,
		inputs:       inputs,
		outputs:      outputs,
		l:            log.New(os.Stdout, logPrefix, 0),
	}
}

func Run(name string, program []int, inputs, outputs chan int) {
	c := new(name, program, inputs, outputs)

	var op opCode
	var param1, param2 int

	for addr := 0; ; {
		instr := c.read(addr)
		c.l.Printf("Instruction: %v <= [%v]\n", instr, addr)
		op = opCode(instr % 100)
		if op == halt {
			close(outputs)
			return
		}

		param1 = c.getParam(addr, 1)

		switch op {
		case add:
			param2 = c.getParam(addr, 2)
			c.l.Printf("ADD: %d + %d", param1, param2)
			c.write(addr+3, param1+param2)
			addr += 4

		case multiply:
			param2 = c.getParam(addr, 2)
			// Parameters that an instruction writes to will never be in immediate mode.
			c.l.Printf("MUL: %d*%d", param1, param2)
			c.write(addr+3, param1*param2)
			addr += 4

		case input:
			c.l.Println("Waiting for INPUT")
			in := <-inputs
			c.l.Printf("Got INPUT: %d", in)
			c.write(addr+1, in)
			addr += 2

		case output:
			c.l.Printf("OUTPUT: %d \n", param1)
			outputs <- param1
			addr += 2

		case jumpIfTrue:
			if param1 != 0 {
				param2 = c.getParam(addr, 2)
				c.l.Printf("JUMP: => %d\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case jumpIfFalse:
			if param1 == 0 {
				param2 = c.getParam(addr, 2)
				c.l.Printf("JUMP: => %d\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case lessThan:
			param2 = c.getParam(addr, 2)
			c.l.Println("LESSTHAN:", param1, param2)
			if param1 < param2 {
				c.write(addr+3, 1)
			} else {
				c.write(addr+3, 0)
			}
			addr += 4

		case equals:
			param2 = c.getParam(addr, 2)
			c.l.Println("EQUALS:", param1, param2)
			if param1 == param2 {
				c.write(addr+3, 1)
			} else {
				c.write(addr+3, 0)
			}
			addr += 4

		case relativeBaseOffset:
			c.relativeBase += param1
			c.l.Printf("RELBASE + %v = %v\n", param1, c.relativeBase)
			addr += 2

		default:
			c.l.Println("Invalid opcode:", op)
			os.Exit(1)
		}
	}
}
