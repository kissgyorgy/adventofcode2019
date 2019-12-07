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
	code := intcode.Load(intcodeFile)
	memory := intcode.Init(code)
	fmt.Println("Memory:", memory)
	intcode.Run(memory, inputVal)
}
