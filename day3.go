package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	up    = "U"
	down  = "D"
	left  = "L"
	right = "R"
)

type Point struct {
	x, y int
}

func getCoordinates(paths []string) map[Point]uint {
	x, y := 0, 0
	var totalSteps uint = 0
	points := make(map[Point]uint, len(paths)*10)
	var stepMeth func()

	for _, path := range paths {
		direction := string(path[0])
		stepsStr, _ := strconv.ParseInt(path[1:], 10, 0)
		steps := int(stepsStr)

		switch direction {
		case up:
			stepMeth = func() { y++ }
		case down:
			stepMeth = func() { y-- }
		case left:
			stepMeth = func() { x-- }
		case right:
			stepMeth = func() { x++ }
		}

		for steps > 0 {
			// skipping first step, bc it's already in there
			// or (0, 0) at start, which doesn't matter
			stepMeth()
			totalSteps++
			steps--
			if _, ok := points[Point{x, y}]; ok {
				// we already has been here, so the number of steps
				// should be the value from the first time
				continue
			} else {
				points[Point{x, y}] = totalSteps
			}
		}
	}

	return points
}

func findCommonPoints(first, second map[Point]uint) map[Point][2]uint {
	common := make(map[Point][2]uint)

	for firstPoint, firstSteps := range first {
		if secondSteps, ok := second[firstPoint]; ok {
			fmt.Println("Found crossing:", firstPoint)
			common[firstPoint] = [2]uint{firstSteps, secondSteps}
		}
	}
	return common
}

func findMinSteps(points map[Point][2]uint) uint {
	var minPoint Point
	var minStepSum uint

	// we take a random element, because maps don't have a notion of "first"
	// and this is the simplest data type we can use for this function,
	// because Go doesn't have a tuple type
	for point, steps := range points {
		minPoint = point
		minStepSum = steps[0] + steps[1]
	}

	for point, steps := range points {
		stepSum := steps[0] + steps[1]
		if stepSum < minStepSum {
			minPoint = point
			minStepSum = stepSum
		}
	}

	fmt.Printf("Minimum steps: %v, with points: %v\n", minStepSum, minPoint)
	return minStepSum
}

func main() {
	content, _ := ioutil.ReadFile("day3-input.txt")
	lines := strings.Split(string(content), "\n")

	wireCoordinates := make([]map[Point]uint, 2, 2)

	for i := 0; i < 2; i++ {
		fmt.Printf("%v. wire ", i+1)
		paths := strings.Split(lines[i], ",")
		wireCoordinates[i] = getCoordinates(paths)
		fmt.Printf("is %v long\n", len(wireCoordinates[i]))
	}

	commonPoints := findCommonPoints(wireCoordinates[0], wireCoordinates[1])
	minStepSum := findMinSteps(commonPoints)
	fmt.Println("Result:", minStepSum)
}
