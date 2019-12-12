package point

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Rad float64

func (p Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}

func (p Point) Minus(p2 Point) Point {
	return Point{p.X - p2.X, p.Y - p2.Y}
}

func DotProduct(p1, p2 Point) float64 {
	return float64(p1.X*p2.X + p1.Y*p2.Y)
}

func (p Point) Length() float64 {
	dotProdSelf := float64(DotProduct(p, p))
	return math.Sqrt(dotProdSelf)
}

func Angle(p1, p2 Point) Rad {
	cosAng := DotProduct(p1, p2) / (p1.Length() * p2.Length())
	return Rad(math.Acos(cosAng))
}

func CrossProduct(p1, p2, p3 Point) int {
	dxc := p1.X - p2.X
	dyc := p1.Y - p2.Y

	dxl := p3.X - p2.X
	dyl := p3.Y - p2.Y

	cross := dxc*dyl - dyc*dxl
	return cross
}

func AreOnTheSameLine(p1, p2, p3 Point) bool {
	return CrossProduct(p1, p2, p3) == 0
}

// https://stackoverflow.com/a/11908158/720077
func (p Point) IsBetweenTwoPoints(p1, p2 Point) bool {
	if !AreOnTheSameLine(p, p1, p2) {
		return false
	}

	dxl := p2.X - p1.X
	dyl := p2.Y - p1.Y

	if math.Abs(float64(dxl)) >= math.Abs(float64(dyl)) {
		if dxl > 0 {
			return p1.X <= p.X && p.X <= p2.X
		} else {
			return p2.X <= p.X && p.X <= p1.X
		}
	} else {
		if dyl > 0 {
			return p1.Y <= p.Y && p.Y <= p2.Y
		} else {
			return p2.Y <= p.Y && p.Y <= p1.Y
		}
	}
}
