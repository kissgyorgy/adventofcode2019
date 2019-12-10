package itertools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		seq      []interface{}
		r        int
		expected [][]interface{}
	}{
		{
			seq:      []interface{}{},
			r:        -1,
			expected: [][]interface{}{[]interface{}{}},
		},
		{
			seq:      []interface{}{42},
			r:        -1,
			expected: [][]interface{}{[]interface{}{42}},
		},
		{
			seq:      []interface{}{42},
			r:        1,
			expected: [][]interface{}{[]interface{}{42}},
		},
		{
			seq: []interface{}{5, 6, 7},
			r:   -1,
			expected: [][]interface{}{
				[]interface{}{5, 6, 7},
				[]interface{}{5, 7, 6},
				[]interface{}{6, 5, 7},
				[]interface{}{6, 7, 5},
				[]interface{}{7, 5, 6},
				[]interface{}{7, 6, 5},
			},
		},
		{
			seq: []interface{}{5, 6, 7},
			r:   3,
			expected: [][]interface{}{
				[]interface{}{5, 6, 7},
				[]interface{}{5, 7, 6},
				[]interface{}{6, 5, 7},
				[]interface{}{6, 7, 5},
				[]interface{}{7, 5, 6},
				[]interface{}{7, 6, 5},
			},
		},
		{
			seq: []interface{}{5, 6, 7},
			r:   2,
			expected: [][]interface{}{
				[]interface{}{5, 6},
				[]interface{}{5, 7},
				[]interface{}{6, 5},
				[]interface{}{6, 7},
				[]interface{}{7, 5},
				[]interface{}{7, 6},
			},
		},
	}
	for _, tt := range tests {
		desc := fmt.Sprintf("%v,%d", tt.seq, tt.r)
		t.Run(desc, func(t *testing.T) {
			allPermutations := make([][]interface{}, 0, len(tt.expected))
			for perm := range Permutations(tt.seq, tt.r) {
				allPermutations = append(allPermutations, perm)
			}
			assert.Equal(t, tt.expected, allPermutations)
		})
	}
}
