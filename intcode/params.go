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
)

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
