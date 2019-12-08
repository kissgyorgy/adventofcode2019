package intcode

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Load(filename string) []string {
	content, _ := ioutil.ReadFile(filename)
	stringContent := strings.TrimSpace(string(content))
	return strings.Split(stringContent, ",")
}

func Init(input []string) []int {
	memory := make([]int, len(input))
	for i, v := range input {
		num, _ := strconv.ParseInt(v, 10, 0)
		memory[i] = int(num)
	}
	return memory
}
