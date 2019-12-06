package main

import "testing"

import "fmt"

import "github.com/stretchr/testify/assert"

func TestGetDigit(t *testing.T) {
	tests := []struct {
		num      int
		ind      int
		expected int
	}{
		{1000, 0, 0},
		{987, 1, 8},
		{3, 0, 3},
		{1101, 2, 1},
		{1101, 3, 1},
	}
	for _, tt := range tests {
		desc := fmt.Sprintf("%d->%d.=%d", tt.num, tt.ind, tt.expected)
		t.Run(desc, func(t *testing.T) {
			res := getNthDigitFromRight(tt.num, tt.ind)
			assert.Equal(t, tt.expected, res)
		})
	}
}
