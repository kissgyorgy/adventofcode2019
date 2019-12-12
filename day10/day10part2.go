package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/kissgyorgy/adventofcode2019/point"
)

const (
	// mapFile = "day10-example3.txt"
	mapFile = "day10-input.txt"
)

var (
	// this comes from solution of part1
	selectedPoint = point.Point{26, 36}
	// selectedPoint = point.Point{3, 4} // example1
	// selectedPoint = point.Point{11, 13} // example3
	yAxisVector = point.Point{0, -1}
)

type pointProps struct {
	point          point.Point // the original point
	relativePoint  point.Point // relative to the Point we subtracted
	angleToY       float64
	relativeLength float64
}

// This will calculate the full degrees between the starting point and the relative point
// so return value will be between 0 and 2π
func calculateFullAngle(start, p2 point.Point) float64 {
	// 0 <= angle <= π, see: https://en.wikipedia.org/wiki/Inverse_trigonometric_functions#arccos
	angle := float64(point.Angle(start, p2))
	fmt.Printf("P: %v Angle: %v\n", p2, angle*180/math.Pi)
	// this is needed so we can have a number between 0 and 2π
	// because rotation starts upside, everything to the left is π+x rad
	// so it will come later when sorted.
	if start.X > p2.X {
		angle = 2*math.Pi - angle
	}
	// we need this, because otherwise the results wouldn't match with example3
	return math.Round(angle*100000) / 100000
}

func main() {
	asteroidMap := loadMap(mapFile)
	printMap(asteroidMap)
	asteroidCoords := convertMapToPoints(asteroidMap)

	asteroidProps := make([]pointProps, 0, len(asteroidCoords)-1)

	for _, p := range asteroidCoords {
		if p == selectedPoint {
			continue
		}
		relativePoint := p.Minus(selectedPoint)
		props := pointProps{
			point:          p,
			relativePoint:  relativePoint,
			angleToY:       calculateFullAngle(yAxisVector, relativePoint),
			relativeLength: relativePoint.Length(),
		}
		asteroidProps = append(asteroidProps, props)
	}

	sort.Slice(asteroidProps, func(i, j int) bool {
		pi, pj := asteroidProps[i], asteroidProps[j]
		if pi.angleToY == pj.angleToY {
			return pi.relativeLength < pj.relativeLength
		}
		return pi.angleToY < pj.angleToY
	})

	fmt.Println("Sorted:")
	for _, p := range asteroidProps {
		fmt.Println(p)
	}

	fmt.Println("Result:")
	nth := 1
	for len(asteroidProps) > 0 {
		var prevAngle float64 = 999
		for i := 0; i < len(asteroidProps); i++ {
			props := asteroidProps[i]
			if prevAngle == props.angleToY {
				continue
			}
			fmt.Printf("%d. %v\n", nth, props)
			nth++
			if i < len(asteroidProps)-1 {
				asteroidProps = append(asteroidProps[:i], asteroidProps[i+1:]...)
				i--
			} else {
				asteroidProps = asteroidProps[:len(asteroidProps)-1]
			}
			prevAngle = props.angleToY
		}
	}
}
