package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/kissgyorgy/adventofcode2019/intcode"
	"github.com/kissgyorgy/adventofcode2019/itertools"
)

const (
	amplifierControllerSoftware = "day7-input.txt"
)

var (
	phaseSettings = []int{0, 1, 2, 3, 4}
	amplifiers    = []string{"A", "B", "C", "D", "E"}
)

func runPhase(phase, program []int, results chan<- int) {
	var out int
	inputSignal := 0
	for i, amp := range amplifiers {
		inputs, outputs := make(chan int, 2), make(chan int)
		inputs <- phase[i]
		inputs <- inputSignal
		go intcode.Run(amp, program, inputs, outputs)
		out = <-outputs
		fmt.Println("Output:", out)
		inputSignal = out
	}
	results <- out
}

func convert(numbers []int) []interface{} {
	ifs := make([]interface{}, len(numbers))
	for i, num := range numbers {
		ifs[i] = interface{}(num)
	}
	return ifs
}

func convertBack(ifs []interface{}) []int {
	numbers := make([]int, len(ifs))
	for i, num := range numbers {
		numbers[i] = int(num)
	}
	return numbers
}

func runSettingPermutations(program, phaseSettings []int, results chan<- int) {
	var wg sync.WaitGroup

	for phase := range itertools.Permutations(convert(phaseSettings), -1) {
		wg.Add(1)
		go func(phase []int) {
			defer wg.Done()
			fmt.Println("Running with phase settings:", phase)
			runPhase(phase, program, results)
		}(convertBack(phase))
	}

	wg.Wait()
	close(results)
}

func collectMaxThrustResults(results <-chan int) int {
	var maxThrust float64 = 0
	for res := range results {
		maxThrust = math.Max(float64(res), maxThrust)
	}
	return int(maxThrust)
}

func main() {
	program := intcode.Load(amplifierControllerSoftware)
	results := make(chan int, 10)
	go runSettingPermutations(program, phaseSettings, results)
	maxThrust := collectMaxThrustResults(results)
	fmt.Println("Result:", maxThrust)
}
