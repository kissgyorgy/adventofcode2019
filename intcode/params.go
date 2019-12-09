package intcode

import (
	"fmt"
	"log"
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

func getParam(l *log.Logger, memory []int, relativeBase, opAddr, nth int) int {
	var val int
	mode := getNthDigitFromRight(memory[opAddr], nth+1)
	param := memory[opAddr+nth]

	switch paramMode(mode) {
	case positionMode:
		val = memory[param]
	case immediateMode:
		val = param
	case relativeMode:
		relativePos := relativeBase + param
		val = memory[relativePos]
	default:
		fmt.Println("Invalid parameter mode:", mode)
		os.Exit(1)
	}
	l.Printf("param%d: %v (addr: %d)\n", nth, val, opAddr+nth)
	return val
}
