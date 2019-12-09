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

		for i, amp := range amplifiers {
			inputs, outputs := make(chan int, 2), make(chan int)
			inputs <- phase[i]
			inputs <- inputSignal
			go intcode.Run(amp, program, inputs, outputs)
			out := <-outputs
			fmt.Println("Output:", out)
			inputSignal = out
			maxThrust = math.Max(float64(out), maxThrust)
			fmt.Println("MaxThrust:", int(maxThrust))
		}
	}
}
