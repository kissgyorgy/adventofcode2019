package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	imageFile = "day8-input.txt"
	width     = 25
	height    = 6
)

var (
	layerSize = width * height
)

func isEmpty(b []byte) bool {
	return strings.TrimSpace(string(b)) == ""
}

func convertToInts(b []byte) []int {
	nums := make([]int, len(b))
	for i, c := range b {
		n, err := strconv.ParseUint(string(c), 10, 8)
		if err != nil {
			panic("Invalid pixel:" + string(c))
		}
		nums[i] = int(n)
	}
	return nums
}

func countDigits(nums []int, digit int) int {
	count := 0
	for _, n := range nums {
		if n == digit {
			count++
		}
	}
	return count
}

func main() {
	var minZeros int = math.MaxInt64
	f, _ := os.Open(imageFile)
	defer f.Close()

	layer := make([]byte, layerSize)
	minPixels := make([]int, layerSize)

	for {
		n, err := io.ReadFull(f, layer)
		if err == io.EOF || (err == io.ErrUnexpectedEOF && isEmpty(layer[:n])) {
			break
		} else if err != nil {
			panic("Could not read layer" + fmt.Sprintf("%v", err))
		}
		pixels := convertToInts(layer)
		zeros := countDigits(pixels, 0)
		if zeros < minZeros {
			minPixels = pixels
			minZeros = zeros
		}
	}

	res := countDigits(minPixels, 1) * countDigits(minPixels, 2)
	fmt.Println("Result:", res)
}
