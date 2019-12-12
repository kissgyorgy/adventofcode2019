package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/kissgyorgy/adventofcode2019/aocimage"
	"github.com/kissgyorgy/adventofcode2019/point"
)

func makeImage(spacecraftSide map[point.Point]myColor) image.Image {
	m := image.NewNRGBA(image.Rect(0, 0, 40, 6))
	for point, col := range spacecraftSide {
		v := uint8(col) * 100
		col := color.RGBA{0, v, v, 255}
		m.Set(point.X, point.Y, col)
	}
	return m
}

func main() {
	// the emergency hull painting robot starting panel" is white, every other panel is still black
	spacecraftSide := paintSpaceShip(white)

	img := makeImage(spacecraftSide)
	scaled := aocimage.Scale(img, 5)
	aocimage.SavePNG(scaled, "plate.png")

	fmt.Println("Result:", len(spacecraftSide))
}
