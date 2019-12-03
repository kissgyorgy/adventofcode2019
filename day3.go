package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func getCoordinates(paths []string) map[Point]bool {
	x, y := 0, 0
	points := make(map[Point]bool, len(paths)*10)
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
			// we can do this because crossing the wire with itself doesn't matter
			points[Point{x, y}] = true
			steps--
		}
	}

	return points
}

func findCommonPoints(first, second map[Point]bool) []Point {
	maxLen := int(math.Max(float64(len(first)), float64(len(second))))
	common := make([]Point, 0, maxLen)

	for firstPoint := range first {
		if second[firstPoint] {
			fmt.Println("Found crossing:", firstPoint)
			common = append(common, firstPoint)
		}
	}
	return common
}

func findMinimumDistance(points []Point) int {
	minPoint := points[0]
	minDist := distanceToZero(minPoint)

	for _, p := range points {
		dist := distanceToZero(p)
		if dist < minDist {
			minPoint = p
			minDist = dist
		}
	}

	fmt.Printf("Minimum distance: %v, with points: %v\n", minDist, minPoint)
	return minDist
}

func distanceToZero(coord Point) int {
	x, y := float64(coord.x), float64(coord.y)
	return int(math.Abs(x) + math.Abs(y))
}

func main() {
	content, _ := ioutil.ReadFile("day3-input.txt")
	lines := strings.Split(string(content), "\n")

	wireCoordinates := make([]map[Point]bool, 2, 2)

	for i := 0; i < 2; i++ {
		fmt.Printf("%v. wire ", i+1)
		paths := strings.Split(lines[i], ",")
		wireCoordinates[i] = getCoordinates(paths)
		fmt.Printf("is %v long\n", len(wireCoordinates[i]))
	}

	commonPoints := findCommonPoints(wireCoordinates[0], wireCoordinates[1])
	minDist := findMinimumDistance(commonPoints)
	fmt.Println("Result:", minDist)
}
