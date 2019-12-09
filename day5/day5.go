package main

import (
	"fmt"

	"github.com/kissgyorgy/adventofcode2019/intcode"
)

const (
	intcodeFile = "day5-input.txt"
	inputVal    = 5
)

func main() {
	program := intcode.Load(intcodeFile)
	fmt.Println("Program:", program)
	inputs, outputs := make(chan int, 1), make(chan int, 1)
	inputs <- inputVal
	intcode.Run("day5", program, inputs, outputs)
}
