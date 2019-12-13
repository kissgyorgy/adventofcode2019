package main

import (
	"fmt"

	"github.com/kissgyorgy/adventofcode2019/intcode"
	"github.com/kissgyorgy/adventofcode2019/point"
)

const (
	arcadeGameFile = "day13-input.txt"
)

type tileId int

const (
	empty  tileId = 0
	wall   tileId = 1
	block  tileId = 2
	paddle tileId = 3
	ball   tileId = 4
)

func main() {
	program := intcode.Load(arcadeGameFile)
	inputs, outputs := make(chan int), make(chan int)
	go intcode.Run("Arcade", program, inputs, outputs)

	blockTiles := make(map[point.Point]bool)
	for {
		x, ok := <-outputs
		if !ok {
			break
		}
		y := <-outputs
		tid := <-outputs
		if tileId(tid) == block {
			blockTiles[point.Point{X: x, Y: y}] = true
		}
	}

	fmt.Println("Result:", len(blockTiles))
}
