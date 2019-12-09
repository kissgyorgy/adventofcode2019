package intcode

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Load(filename string) []int {
	content, _ := ioutil.ReadFile(filename)
	stringContent := strings.TrimSpace(string(content))
	code := strings.Split(stringContent, ",")

	program := make([]int, len(code))
	for i, v := range code {
		num, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			msg := fmt.Sprintf("Invalid Intcode instruction: %v", err)
			panic(msg)
		}
		program[i] = int(num)
	}
	return program
}
