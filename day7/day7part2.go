package main

import (
	"fmt"
	"math"

	"github.com/kissgyorgy/adventofcode2019/intcode"
)

const (
	amplifierControllerSoftware = "day7-input.txt"
)

var (
	phaseSettings = []int{5, 6, 7, 8, 9}
	amplifiers    = []string{"A", "B", "C", "D", "E"}
)

func copyMem(from []int) []int {
	newMemory := make([]int, len(from))
	copy(newMemory, from)
	return newMemory
}

func getThrust(initMemory []int, currentSettings []int) int {
	var outputs chan int
	// buffered because we need to put in the phase settings and the initial 0
	// even before anything happens
	outputsE := make(chan int, 2)
	outputsE <- currentSettings[0]
	outputsE <- 0
	prevOutputs := outputsE

	for i, amp := range amplifiers[:len(amplifiers)-1] {
		outputs = make(chan int, 1)
		outputs <- currentSettings[i+1]
		go intcode.Run(amp, copyMem(initMemory), prevOutputs, outputs)
		prevOutputs = outputs
	}

	// don't run E in goroutine, wait for it to finish
	intcode.Run("E", copyMem(initMemory), prevOutputs, outputsE)

	out := <-outputsE
	fmt.Printf("Thruster output: %v\n\n", out)
	return out
}

func main() {
	code := intcode.Load(amplifierControllerSoftware)
	initMemory := intcode.Init(code)

	var maxThrust float64 = 0

	for currentSettings := range IterPermutations(phaseSettings, -1) {
		thrust := getThrust(initMemory, currentSettings)
		maxThrust = math.Max(float64(thrust), maxThrust)
	}
	fmt.Printf("MAX thrust: %d\n", int(maxThrust))
}
