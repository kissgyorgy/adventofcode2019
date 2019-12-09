package intcode

import (
	"fmt"
	"math"
	"os"
)

type paramMode int

const (
	positionMode  paramMode = 0
	immediateMode paramMode = 1
	relativeMode  paramMode = 2
)

// get nth digit of number counted from rigth to left
func getNthDigitFromRight(num, ind int) int {
	return num / int(math.Pow10(ind)) % 10
}

func (c *computer) read(addr int) int {
	if addr > len(c.memory)-1 {
		// we don't even need to expand memory at this point
		return 0
	}
	return c.memory[addr]
}

func (c *computer) write(addr, value int) {
	// Parameters that an instruction writes to will never be in immediate mode.
	pos := c.read(addr)
	c.l.Printf("    %d => [%d]", value, pos)
	if pos > len(c.memory)-1 {
		needed := pos + 1 - len(c.memory)
		extraMem := make([]int, needed)
		c.memory = append(c.memory, extraMem...)
	}
	c.memory[pos] = value
}

func (c *computer) getParam(opAddr, nth int) int {
	var val int
	mode := getNthDigitFromRight(c.read(opAddr), nth+1)
	param := c.read(opAddr + nth)

	switch paramMode(mode) {
	case positionMode:
		val = c.read(param)
	case immediateMode:
		val = param
	case relativeMode:
		relativePos := c.relativeBase + param
		val = c.read(relativePos)
	default:
		fmt.Println("Invalid parameter mode:", mode)
		os.Exit(1)
	}
	c.l.Printf("param%d: %v <= [%d]\n", nth, val, opAddr+nth)
	return val
}
