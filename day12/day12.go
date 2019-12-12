package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/kissgyorgy/adventofcode2019/point"
)

const (
	moonPositionsFile = "day12-input.txt"
	steps             = 1000
)

var moonNames = []string{"Io", "Europa", "Ganymede", "Callisto"}

type Moon struct {
	position point.Point3D
	velocity point.Point3D
}

func parseParts(line string) (int, int, int) {
	parts := strings.Split(line, ", ")
	xStr := parts[0][len("<x="):]
	// ignore error checking, the input is well formatted
	x, _ := strconv.ParseInt(xStr, 10, 0)

	yStr := parts[1][len("y="):]
	y, _ := strconv.ParseInt(yStr, 10, 0)

	zStr := parts[2][len("z=") : len(parts[2])-1]
	z, _ := strconv.ParseInt(zStr, 10, 0)

	return int(x), int(y), int(z)
}

func readMoonPositions(filename string) map[string]*Moon {
	f, _ := os.Open(moonPositionsFile)
	defer f.Close()

	fmt.Println("Loading moons:")
	moons := make(map[string]*Moon)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		x, y, z := parseParts(line)
		position := point.Point3D{X: x, Y: y, Z: z}
		velocity := point.Point3D{X: 0, Y: 0, Z: 0}
		moon := &Moon{position: position, velocity: velocity}
		moons[moonNames[i]] = moon
		fmt.Printf("%d. %s: %v\n", i+1, moonNames[i], *moon)
		i++
	}
	return moons
}

func convert(moonNames []string) []interface{} {
	ifs := make([]interface{}, len(moonNames))
	for i, moonName := range moonNames {
		ifs[i] = interface{}(moonName)
	}
	return ifs
}

func getGravities(first, second int) (int, int) {
	if first < second {
		return +1, -1
	} else if second < first {
		return -1, +1
	} else {
		return 0, 0
	}
}

func applyGravity(moons map[string]*Moon) map[string]point.Point3D {
	velocities := make(map[string]*point.Point3D)

	// we need to increase velocities only after every pair has been compared
	for _, moonName := range moonNames {
		velocities[moonName] = &moons[moonName].velocity
	}

	for i := 0; i < len(moonNames)-1; i++ {
		for j := i + 1; j < len(moonNames); j++ {
			moon1Name, moon2Name := moonNames[i], moonNames[j]
			moon1, moon2 := moons[moon1Name], moons[moon2Name]
			velocity1 := velocities[moon1Name]
			velocity2 := velocities[moon2Name]

			pos1Diff, pos2Diff := getGravities(moon1.position.X, moon2.position.X)
			velocity1.X += pos1Diff
			velocity2.X += pos2Diff

			pos1Diff, pos2Diff = getGravities(moon1.position.Y, moon2.position.Y)
			velocity1.Y += pos1Diff
			velocity2.Y += pos2Diff

			pos1Diff, pos2Diff = getGravities(moon1.position.Z, moon2.position.Z)
			velocity1.Z += pos1Diff
			velocity2.Z += pos2Diff
		}
	}

	// make velocities immutable
	results := make(map[string]point.Point3D)
	for name, vel := range velocities {
		results[name] = *vel
	}
	return results
}

func applyVelocity(moons map[string]*Moon, velocities map[string]point.Point3D) {
	for _, moonName := range moonNames {
		moon := moons[moonName]
		velocity := velocities[moonName]
		moon.position = point.Add3D(moon.position, velocity)
		moon.velocity = velocity
	}
}

func printMoons(moons map[string]*Moon) {
	for _, moonName := range moonNames {
		moon := moons[moonName]
		fmt.Printf("pos=%s, vel=%s\n", moon.position, moon.velocity)
	}
}

func calculateTotalEnergy(moons map[string]*Moon) int {
	totalEnergy := 0
	for _, moon := range moons {
		potentialEnergy := math.Abs(float64(moon.position.X)) + math.Abs(float64(moon.position.Y)) + math.Abs(float64(moon.position.Z))
		kineticEnergy := math.Abs(float64(moon.velocity.X)) + math.Abs(float64(moon.velocity.Y)) + math.Abs(float64(moon.velocity.Z))
		subTotal := int(potentialEnergy * kineticEnergy)
		totalEnergy += subTotal
		fmt.Printf("potential: %3v   kinetic: %3v   subtotal: %v\n", potentialEnergy, kineticEnergy, subTotal)
	}
	fmt.Println("Total:", totalEnergy)
	return totalEnergy
}

func main() {
	moons := readMoonPositions(moonPositionsFile)
	fmt.Println()

	for i := 0; i < steps; i++ {
		velocities := applyGravity(moons)
		applyVelocity(moons, velocities)

		fmt.Printf("After %d steps:\n", i+1)
		printMoons(moons)
		fmt.Println()

		fmt.Printf("Energy after %d steps:\n", i+1)
		calculateTotalEnergy(moons)
	}
}
