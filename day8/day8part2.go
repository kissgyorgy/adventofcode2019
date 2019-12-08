package main

import (
	"bytes"
	"fmt"
	"image"
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
	scale     = 20
)

type color int

const (
	black       color = 0
	white       color = 1
	transparent color = 2
)

var (
	layerSize = width * height
)

func isEmpty(b []byte) bool {
	return strings.TrimSpace(string(b)) == ""
}

func convertToPixels(b []byte) []color {
	pixels := make([]color, len(b))
	for i, c := range b {
		n, err := strconv.ParseUint(string(c), 10, 8)
		if err != nil {
			panic("Invalid pixel:" + string(c))
		}
		pixels[i] = color(n)
	}
	return pixels
}

func readLayers(imageFile string) [][]color {
	f, _ := os.Open(imageFile)
	defer f.Close()

	layers := make([][]color, 0, 4)
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

func mergeLayers(layers [][]color) []color {
	image := make([]color, layerSize)
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

func printImage(image []color, width, height int) {
	for i, pix := range image {
		if i%width == 0 && i != 0 {
			fmt.Println()
		}

		if pix == black {
			fmt.Printf(" ")
		} else {
			fmt.Printf("█")
		}
	}
	fmt.Println()
}

func makeImage(img []color) image.Image {
	const (
		dx = width
		dy = height
	)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			ind := y*dx + x
			v := uint8(img[ind]) * 150
			i := y*m.Stride + x*4

			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
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
	printImage(finalLayer, width, height)

	img := makeImage(finalLayer)
	scaled := scaleImage(img, scale)
	savePNG(scaled, fmt.Sprintf("image-%dx.png", scale))
}
