package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateFuel(mass int) int {
	fuelNeeded := int(math.Floor(float64(mass/3))) - 2
	if fuelNeeded <= 0 {
		return 0
	}
	fmt.Println("Need more fuel:", fuelNeeded)
	return fuelNeeded + calculateFuel(fuelNeeded)
}

func main() {
	f, _ := os.Open("day1-input.txt")
	defer f.Close()

	sum, totalFuelNeededForMass := 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		massStr := scanner.Text()
		mass, _ := strconv.ParseInt(massStr, 10, 64)
		fmt.Println("Got mass:", mass)
		totalFuelNeededForMass = calculateFuel(int(mass))
		fmt.Println("Total fuel needed for mass:", totalFuelNeededForMass)
		sum += totalFuelNeededForMass
	}

	fmt.Println("Summary:", sum)
}
