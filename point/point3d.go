package point

import "fmt"

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("<x=%3d, y=%3d, z=%3d>", p.X, p.Y, p.Z)
}

func Add3D(p1, p2 Point3D) Point3D {
	return Point3D{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
		Z: p1.Z + p2.Z,
	}
}
