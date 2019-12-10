package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.x, p.y)
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

}
