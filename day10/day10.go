package main

import (
	"fmt"

	"github.com/kissgyorgy/adventofcode2019/point"
)

const (
	mapFile = "day10-input.txt"
)

func isAPointBetween(p1 point.Point, points []point.Point, p2 point.Point) bool {
	for _, middle := range points {
		if middle == p1 || middle == p2 {
			continue
		}
		if middle.IsBetweenTwoPoints(p1, p2) {
			return true
		}
	}
	return false
}

func main() {
	asteroidMap := loadMap(mapFile)
	asteroidCoords := convertMapToPoints(asteroidMap)
	fmt.Println(asteroidCoords)

	detectableAstroids := make(map[point.Point]int)

	for i, p1 := range asteroidCoords {
		for j := i + 1; j < len(asteroidCoords); j++ {
			p2 := asteroidCoords[j]
			if p1 == p2 {
				continue
			}
			if !isAPointBetween(p1, asteroidCoords, p2) {
				detectableAstroids[p1]++
				detectableAstroids[p2]++
			}
		}
	}

	maxAsteroids := 0
	for _, p := range asteroidCoords {
		count := detectableAstroids[p]
		fmt.Println(p, count)
		if count > maxAsteroids {
			maxAsteroids = count
		}
	}
	fmt.Println("Result:", maxAsteroids)
}
