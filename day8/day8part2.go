package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	imageFile = "day8-input.txt"
	width     = 25
	height    = 6
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
			fmt.Printf("*")
		}
	}
	fmt.Println()
}

func main() {
	layers := readLayers(imageFile)
	image := mergeLayers(layers)
	printImage(image, width, height)
}
