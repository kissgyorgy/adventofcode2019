package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	reactionsFile = "day14-input.txt"
)

type chemical string

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

func main() {
	content, _ := ioutil.ReadFile(reactionsFile)
	reactions := parseReactions(content)
	for _, r := range reactions {
		fmt.Println(r)
	}
}
