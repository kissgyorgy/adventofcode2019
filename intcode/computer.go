package intcode

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const Silent = true

type computer struct {
	memory       []int
	relativeBase int
	inputs       <-chan int
	outputs      chan<- int
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

func new(name string, program []int, inputs <-chan int, outputs chan<- int, quiet bool) *computer {
	logPrefix := fmt.Sprintf("[%s] ", name)
	mem := make([]int, len(program))
	copy(mem, program)

	logger := log.New(os.Stdout, logPrefix, 0)
	if quiet {
		logger.SetOutput(ioutil.Discard)
	}

	return &computer{
		memory:       mem,
		relativeBase: 0,
		inputs:       inputs,
		outputs:      outputs,
		l:            logger,
	}
}

func getQuiet(quiet []bool) bool {
	q := false
	if len(quiet) == 1 {
		q = quiet[0]
	} else if len(quiet) > 1 {
		panic("Invalid function call")
	}
	return q
}

func Run(name string, program []int, inputs <-chan int, outputs chan<- int, quiet ...bool) {
	c := new(name, program, inputs, outputs, getQuiet(quiet))

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
		param2 = c.getParam(addr, 2)

		switch op {
		case add:
			c.l.Printf("ADD: %d + %d", param1, param2)
			c.write(addr, 3, param1+param2)
			addr += 4

		case multiply:
			// Parameters that an instruction writes to will never be in immediate mode.
			c.l.Printf("MUL: %d*%d", param1, param2)
			c.write(addr, 3, param1*param2)
			addr += 4

		case input:
			c.l.Println("Waiting for INPUT")
			in := <-inputs
			c.l.Printf("Got INPUT: %d", in)
			c.write(addr, 1, in)
			addr += 2

		case output:
			c.l.Printf("OUTPUT: %d \n", param1)
			outputs <- param1
			addr += 2

		case jumpIfTrue:
			if param1 != 0 {
				c.l.Printf("JUMP: => %d\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case jumpIfFalse:
			if param1 == 0 {
				c.l.Printf("JUMP: => %d\n", param2)
				addr = param2
			} else {
				addr += 3
			}

		case lessThan:
			c.l.Println("LESSTHAN:", param1, param2)
			if param1 < param2 {
				c.write(addr, 3, 1)
			} else {
				c.write(addr, 3, 0)
			}
			addr += 4

		case equals:
			c.l.Println("EQUALS:", param1, param2)
			if param1 == param2 {
				c.write(addr, 3, 1)
			} else {
				c.write(addr, 3, 0)
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
