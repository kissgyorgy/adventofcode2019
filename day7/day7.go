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
	amplifiers    = []string{"A", "B", "C", "D", "E"}
)

func main() {
	program := intcode.Load(amplifierControllerSoftware)
	memory := make([]int, len(program))

	var maxThrust float64 = 0

	for phase := range IterPermutations(phaseSettings, -1) {
		fmt.Println("Phase settings:", phase)
		copy(memory, program)
		inputSignal := 0
		for i, _ := range amplifiers {
			outputs := intcode.Run(program, phase[i], inputSignal)
			fmt.Println("Outputs:", outputs)
			inputSignal = outputs[0]
			maxThrust = math.Max(float64(outputs[0]), maxThrust)
			fmt.Println("MaxThrust:", int(maxThrust))
		}
	}
}
