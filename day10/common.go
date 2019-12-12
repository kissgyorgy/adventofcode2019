package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kissgyorgy/adventofcode2019/point"
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

func printMap(asteroidMap [][]byte) {
	fmt.Print("   ")
	for i := range asteroidMap[0] {
		fmt.Printf("%3d", i)
	}
	fmt.Println()
	for i, l := range asteroidMap {
		fmt.Printf("%3d", i)
		for _, cell := range l {
			fmt.Printf("%3s", string(cell))
		}
		fmt.Println()
	}
}
