package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/kissgyorgy/adventofcode2019/aocimage"
)

const (
	imageFile = "day8-input.txt"
	width     = 25
	height    = 6
	scale     = 2
)

type myColor int

const (
	black       myColor = 0
	white       myColor = 1
	transparent myColor = 2
)

var (
	layerSize = width * height
)

func isEmpty(b []byte) bool {
	return strings.TrimSpace(string(b)) == ""
}

func convertToPixels(b []byte) []myColor {
	pixels := make([]myColor, len(b))
	for i, c := range b {
		n, err := strconv.ParseUint(string(c), 10, 8)
		if err != nil {
			panic("Invalid pixel:" + string(c))
		}
		pixels[i] = myColor(n)
	}
	return pixels
}

func readLayers(imageFile string) [][]myColor {
	f, _ := os.Open(imageFile)
	defer f.Close()

	layers := make([][]myColor, 0, 4)
	layer := make([]byte, layerSize)

	i := -1
	for {
		i++
		n, err := io.ReadFull(f, layer)
		if err == io.EOF || (err == io.ErrUnexpectedEOF && isEmpty(layer[:n])) {
			break
		} else if err != nil {
			panic("Could not read layer" + fmt.Sprintf("%v", err))
		}
		pixels := convertToPixels(layer)
		layers = append(layers, pixels)
	}
	return layers
}

func mergeLayers(layers [][]myColor) []myColor {
	image := make([]myColor, layerSize)
	for i := 0; i < layerSize; i++ {
		for _, layer := range layers {
			pix := layer[i]
			if pix == black || pix == white {
				image[i] = pix
				break
			}
		}
	}
	return image
}

func makeImage(img []myColor, width, height int) image.Image {
	m := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := y*width + x
			v := uint8(img[i]) * 100
			m.Set(x, y, color.RGBA{0, v, v, 255})
		}
	}
	return m
}

func main() {
	layers := readLayers(imageFile)
	finalLayer := mergeLayers(layers)
	img := makeImage(finalLayer, width, height)
	aocimage.Print(img)
	scaled := aocimage.Scale(img, scale)
	aocimage.SavePNG(scaled, fmt.Sprintf("image-%dx.png", scale))
}
