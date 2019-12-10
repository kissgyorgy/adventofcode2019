package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	mapFile = "day10-example1.txt"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.x, p.y)
}

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

func convertToPoints(mapLines [][]byte) []Point {
	points := make([]Point, 0, 100)
	for y, line := range loadMap(mapFile) {
		for x, char := range line {
			if char == '#' {
				p := Point{x, y}
				points = append(points, p)
			}
		}
	}
	return points
}

func crossProduct(p1, p2, p3 Point) int {
	dxc := p1.x - p2.x
	dyc := p1.y - p2.y

	dxl := p3.x - p2.x
	dyl := p3.y - p2.y

	cross := dxc*dyl - dyc*dxl
	return cross
}

func areOnTheSameLine(p1, p2, p3 Point) bool {
	return crossProduct(p1, p2, p3) == 0
}

// https://stackoverflow.com/a/11908158/720077
func (p Point) isBetweenTwoPoints(p1, p2 Point) bool {
	if !areOnTheSameLine(p, p1, p2) {
		return false
	}

	dxl := p2.x - p1.x
	dyl := p2.y - p1.y

	if math.Abs(float64(dxl)) >= math.Abs(float64(dyl)) {
		if dxl > 0 {
			return p1.x <= p.x && p.x <= p2.x
		} else {
			return p2.x <= p.x && p.x <= p1.x
		}
	} else {
		if dyl > 0 {
			return p1.y <= p.y && p.y <= p2.y
		} else {
			return p2.y <= p.y && p.y <= p1.y
		}
	}
}

func main() {
	asteroidMap := loadMap(mapFile)
	for _, line := range asteroidMap {
		fmt.Println(string(line))
	}
	fmt.Println(convertToPoints(asteroidMap))
}
