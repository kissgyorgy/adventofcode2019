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
	phaseSettings = []int{0, 1, 2, 3, 4}
)

func main() {
	code := intcode.Load(amplifierControllerSoftware)
	initMemory := intcode.Init(code)
	memory := make([]int, len(initMemory), len(initMemory))

	var maxThrust float64 = 0

	for phase := range IterPermutations(phaseSettings, -1) {
		fmt.Println("Phase settings:", phase)
		copy(memory, initMemory)
		inputSignal := 0
		for i := 0; i < 5; i++ {
			outputs := intcode.Run(initMemory, phase[i], inputSignal)
			fmt.Println("Outputs:", outputs)
			inputSignal = outputs[0]
			maxThrust = math.Max(float64(outputs[0]), maxThrust)
			fmt.Println("MaxThrust:", int(maxThrust))
		}
	}
}
