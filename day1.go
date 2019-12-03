package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}

func main() {
	f, _ := os.Open("day1-input.txt")
	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		massStr := scanner.Text()
		mass, _ := strconv.ParseInt(massStr, 10, 64)
		sum += calculateFuel(int(mass))
	}

	fmt.Println("Summary:", sum)
}
