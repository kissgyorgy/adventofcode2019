package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func loadObjects(filename string) map[string][]string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Cannot load input data")
	}

	objects := make(map[string][]string)

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		relation := strings.Split(line, ")")
		name, orbiter := relation[0], relation[1]
		objects[name] = append(objects[name], orbiter)
	}
	return objects
}

func countOrbitsByLevel(objects map[string][]string, orbitee string, level int) int {
	orbiters, ok := objects[orbitee]
	if !ok {
		// an object which has no orbiter
		return level
	}

	subLevels := 0
	for _, orbiter := range orbiters {
		subLevels += countOrbitsByLevel(objects, orbiter, level+1)
	}
	return level + subLevels
}

func main() {
	objects := loadObjects("day6-input.txt")
	// we start by looking up the root object at level 0
	// (it is not an orbiter, so doesn't count)
	res := countOrbitsByLevel(objects, "COM", 0)
	fmt.Println("Result:", res)
}
