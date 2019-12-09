package main

import (
	"fmt"

	"github.com/kissgyorgy/adventofcode2019/intcode"
)

const (
	boostCode       = "day9-input.txt"
	testMode        = 1
	sensorBoostMode = 2
)

func main() {
	program := intcode.Load(boostCode)

	inputs, outputs := make(chan int, 1), make(chan int)
	// inputs <- testMode
	inputs <- sensorBoostMode
	go intcode.Run("day9", program, inputs, outputs)

	fullOut := make([]int, 0, 10)
	for out := range outputs {
		fullOut = append(fullOut, out)
	}
	fmt.Print("Out:", fullOut)
}
