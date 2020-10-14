package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	// reactionsFile = "day14-input.txt"
	reactionsFile = "example1.txt"
)

type chemical string

const (
	ORE  = chemical("ORE")
	FUEL = chemical("FUEL")
)

type chemIO struct {
	chem chemical
	qty  int
}

type reaction struct {
	inputs []chemIO
	output chemIO
}

type nanofactory struct{}

func parseParts(partlist string) []chemIO {
	chemIOs := make([]chemIO, 0, 10)

	chems := strings.Split(partlist, ", ")
	for _, ch := range chems {
		parts := strings.Split(ch, " ")
		qty, _ := strconv.ParseInt(parts[0], 10, 0)
		chem := chemIO{
			chem: chemical(parts[1]),
			qty:  int(qty),
		}
		chemIOs = append(chemIOs, chem)
	}

	return chemIOs
}

func parseReactions(content []byte) []reaction {
	reactions := make([]reaction, 0, 10)
	reader := bytes.NewReader(content)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		r := reaction{
			inputs: parseParts(parts[0]),
			output: parseParts(parts[1])[0],
		}
		reactions = append(reactions, r)
	}

	return reactions
}

func searchInputs(reactions []reaction, chem chemical, qty int) int {
	var inputs []chemIO
	var mul, outQty int

	fmt.Println("Start searching", chem)
	for _, r := range reactions {
		fmt.Println("Current reaction:", r)
		if r.output.chem == chem {
			inputs = r.inputs
			outQty = r.output.qty
			mul = int(math.Ceil(float64(qty / outQty)))
			fmt.Println("Found inputs for chem:", inputs)
			break
		}
	}
	fmt.Println("Inputs:", inputs)

	if inputs[0].chem == ORE {
		fmt.Println("Found ORE:", inputs[0].qty)
		return inputs[0].qty
	}

	var sum int
	for _, ch := range inputs {
		fmt.Println("Searching for", ch.chem)
		sum += searchInputs(reactions, ch.chem, mul*ch.qty)
		fmt.Println("Subtotal:", sum)
	}

	return sum
}

func main() {
	content, _ := ioutil.ReadFile(reactionsFile)
	reactions := parseReactions(content)
	for _, r := range reactions {
		fmt.Println(r.inputs, "-->", r.output)
	}
	total := searchInputs(reactions, FUEL, 1)
	fmt.Println("Result:", total)
}
