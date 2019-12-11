package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kissgyorgy/adventofcode2019/point"
)

const (
	mapFile = "day10-input.txt"
)

func loadMap(mapFile string) [][]byte {
	file, _ := os.Open(mapFile)
	defer file.Close()

	mapLines := make([][]byte, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mapLines = append(mapLines, scanner.Bytes())
	}
	return mapLines
}

func convertMapToPoints(mapLines [][]byte) []point.Point {
	points := make([]point.Point, 0, 100)
	for y, line := range loadMap(mapFile) {
		for x, char := range line {
			if char == '#' {
				p := point.Point{X: x, Y: y}
				points = append(points, p)
			}
		}
	}
	return points
}

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
