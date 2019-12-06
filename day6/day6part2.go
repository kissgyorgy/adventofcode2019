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

func findSubTree(objects map[string][]string, subTree []string) []string {
	target := subTree[0]
	for parent, orbiters := range objects {
		for _, orbiter := range orbiters {
			if orbiter == target {
				subTree = append([]string{parent}, subTree...)
				subTree = findSubTree(objects, subTree)
				break
			}
		}
	}
	return subTree
}

func findCommonParentLevel(subTree1, subTree2 []string) (string, int) {
	for level, obj := range subTree1 {
		if obj != subTree2[level] {
			commonParent := subTree1[level-1]
			commonLevel := level - 1
			return commonParent, commonLevel
		}
	}
	// the two trees are exactly the same
	return subTree1[0], 0
}

func main() {
	objects := loadObjects("day6-input.txt")
	for name, orbiters := range objects {
		fmt.Printf("%s --> %s\n", name, orbiters)
	}

	subTree1 := findSubTree(objects, []string{"SAN"})
	fmt.Println("SAN subtree:", subTree1)

	subTree2 := findSubTree(objects, []string{"YOU"})
	fmt.Println("YOU subtree:", subTree2)

	commonName, commonLevel := findCommonParentLevel(subTree1, subTree2)
	fmt.Println("Common parent level:", commonName, commonLevel)

	tree1Level := len(subTree1) - 1
	tree2Level := len(subTree2) - 1
	res := tree1Level + tree2Level - 2*commonLevel - 2
	fmt.Println("Result:", res)
}
