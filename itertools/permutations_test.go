package itertools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		seq      []int
		r        int
		expected [][]int
	}{
		{
			seq:      []int{},
			r:        -1,
			expected: [][]int{[]int{}},
		},
		{
			seq:      []int{42},
			r:        -1,
			expected: [][]int{[]int{42}},
		},
		{
			seq:      []int{42},
			r:        1,
			expected: [][]int{[]int{42}},
		},
		{
			seq: []int{5, 6, 7},
			r:   -1,
			expected: [][]int{
				[]int{5, 6, 7},
				[]int{5, 7, 6},
				[]int{6, 5, 7},
				[]int{6, 7, 5},
				[]int{7, 5, 6},
				[]int{7, 6, 5},
			},
		},
		{
			seq: []int{5, 6, 7},
			r:   3,
			expected: [][]int{
				[]int{5, 6, 7},
				[]int{5, 7, 6},
				[]int{6, 5, 7},
				[]int{6, 7, 5},
				[]int{7, 5, 6},
				[]int{7, 6, 5},
			},
		},
		{
			seq: []int{5, 6, 7},
			r:   2,
			expected: [][]int{
				[]int{5, 6},
				[]int{5, 7},
				[]int{6, 5},
				[]int{6, 7},
				[]int{7, 5},
				[]int{7, 6},
			},
		},
	}
	for _, tt := range tests {
		desc := fmt.Sprintf("%v,%d", tt.seq, tt.r)
		t.Run(desc, func(t *testing.T) {
			allPermutations := make([][]int, 0, len(tt.expected))
			for perm := range Permutations(tt.seq, tt.r) {
				allPermutations = append(allPermutations, perm)
			}
			assert.Equal(t, tt.expected, allPermutations)
		})
	}
}
