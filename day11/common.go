package main

import (
	"github.com/kissgyorgy/adventofcode2019/intcode"
	"github.com/kissgyorgy/adventofcode2019/point"
)

const (
	robotProgram = "day11-input.txt"
)

type myColor int

const (
	black myColor = 0
	white myColor = 1
)

type direction int

const (
	left  direction = 0
	right direction = 1
	up    direction = 3
	down  direction = 4
)

func paintSpaceShip(startingPoint myColor) map[point.Point]myColor {
	program := intcode.Load(robotProgram)
	inputs, outputs := make(chan int, 1), make(chan int)
	go intcode.Run("painting robot", program, inputs, outputs)

	spaceCraftSide := make(map[point.Point]myColor)

	currentDirection := up
	currentPoint := point.Point{X: 0, Y: 0}
	spaceCraftSide[currentPoint] = startingPoint

	for {
		if panelColor, ok := spaceCraftSide[currentPoint]; ok {
			// send the last painted (current) color
			inputs <- int(panelColor)
		} else {
			// every panel is black initially
			inputs <- int(black)
		}

		paintColorInt, ok := <-outputs
		if !ok {
			break
		}

		// we need to know the fact "at least once" painted on white
		// we could miss this information if later one panel is overpainted
		spaceCraftSide[currentPoint] = myColor(paintColorInt)

		nextDirection := direction(<-outputs)
		if nextDirection == left {
			switch currentDirection {
			case up:
				currentPoint.X -= 1
				currentDirection = left
			case down:
				currentPoint.X += 1
				currentDirection = right
			case left:
				currentPoint.Y += 1
				currentDirection = down
			case right:
				currentPoint.Y -= 1
				currentDirection = up
			}
		} else if nextDirection == right {
			switch currentDirection {
			case up:
				currentPoint.X += 1
				currentDirection = right
			case down:
				currentPoint.X -= 1
				currentDirection = left
			case left:
				currentPoint.Y -= 1
				currentDirection = up
			case right:
				currentPoint.Y += 1
				currentDirection = down
			}
		}
	}
	return spaceCraftSide
}
