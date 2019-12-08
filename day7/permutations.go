package main

import (
	"errors"
)

func IterPermutations(numbers []int, r int) <-chan []int {
	if r > len(numbers) {
		err := errors.New("r cannot be bigger than the length of numbers")
		panic(err)
	}

	ch := make(chan []int)
	go func() {
		defer close(ch)
		permutate(ch, numbers, r)
	}()
	return ch
}

// an implementation similar to Python standard library itertools.permutations:
// https://docs.python.org/3.8/library/itertools.html#itertools.permutations
func permutate(ch chan []int, numbers []int, r int) {
	n := len(numbers)

	if r < 0 {
		r = n
	}

	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := 0; i < r; i++ {
		cycles[i] = n - i
	}

	nextPerm := func() []int {
		perm := make([]int, r)
		for i, ind := range indices[:r] {
			perm[i] = numbers[ind]
		}
		return perm
	}

	ch <- nextPerm()

	if n < 2 {
		return
	}

	var tmp []int
	var j int

	for i := r - 1; i > -1; i-- {
		cycles[i] -= 1
		if cycles[i] == 0 {
			tmp = append(indices[i+1:], indices[i])
			indices = append(indices[:i], tmp...)
			cycles[i] = n - i
		} else {
			j = len(indices) - cycles[i]
			indices[i], indices[j] = indices[j], indices[i]
			ch <- nextPerm()
			i = r // start over the cycle
			// i-- will apply, so i will be r-1 at the start of the next cycle
		}
	}
}
