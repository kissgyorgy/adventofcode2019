package itertools

import (
	"errors"
)

func Permutations(elements []interface{}, r int) <-chan []interface{} {
	if r > len(elements) {
		err := errors.New("r cannot be bigger than the length of numbers")
		panic(err)
	}

	ch := make(chan []interface{})
	go func() {
		defer close(ch)
		permutate(ch, elements, r)
	}()
	return ch
}

// an implementation similar to Python standard library itertools.permutations:
// https://docs.python.org/3.8/library/itertools.html#itertools.permutations
func permutate(ch chan<- []interface{}, elements []interface{}, r int) {
	n := len(elements)

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

	nextPerm := func() []interface{} {
		perm := make([]interface{}, r)
		for i, ind := range indices[:r] {
			perm[i] = elements[ind]
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
