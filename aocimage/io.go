package aocimage

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
)

func Print(img image.Image) {
	bounds := img.Bounds()
	pi := image.NewPaletted(bounds, []color.Color{
		color.Gray{Y: 255},
		color.Gray{Y: 160},
		color.Gray{Y: 70},
		color.Gray{Y: 35},
		color.Gray{Y: 0},
	})
	draw.FloydSteinberg.Draw(pi, bounds, img, image.ZP)
	shade := []string{" ", "░", "▒", "▓", "█"}

	for i, p := range pi.Pix {
		fmt.Print(shade[p])
		if (i+1)%bounds.Max.X == 0 {
			fmt.Println()
		}
	}
}

func SavePNG(img image.Image, filename string) {
	fmt.Println("Saving image:", filename)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, buf.Bytes(), 0644)
}
