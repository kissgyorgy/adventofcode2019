package aocimage

import (
	"image"

	"golang.org/x/image/draw"
)

func Scale(src image.Image, scale int) image.Image {
	sbo := src.Bounds()
	dr := image.Rect(0, 0, sbo.Max.X*scale, sbo.Max.Y*scale)
	dst := image.NewRGBA(dr)
	draw.NearestNeighbor.Scale(dst, dr, src, sbo, draw.Over, nil)
	return dst
}
