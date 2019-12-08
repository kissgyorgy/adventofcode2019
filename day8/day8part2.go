package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
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

func printImage(img image.Image, width int) {
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
		if (i+1)%width == 0 {
			fmt.Println()
		}
	}
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

func scaleImage(src image.Image, scale int) image.Image {
	sbo := src.Bounds()
	dr := image.Rect(0, 0, sbo.Max.X*scale, sbo.Max.Y*scale)
	dst := image.NewRGBA(dr)
	draw.NearestNeighbor.Scale(dst, dr, src, sbo, draw.Over, nil)
	return dst
}

func savePNG(img image.Image, filename string) {
	fmt.Println("Saving image:", filename)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, buf.Bytes(), 0644)
}

func main() {
	layers := readLayers(imageFile)
	finalLayer := mergeLayers(layers)
	img := makeImage(finalLayer, width, height)
	printImage(img, width)
	scaled := scaleImage(img, scale)
	savePNG(scaled, fmt.Sprintf("image-%dx.png", scale))
}
