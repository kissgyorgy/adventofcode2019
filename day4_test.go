package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitsNeverDecrease(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{0, true},
		{100, false},
		{102, false},
		{111, true},
		{112, true},
		{123, true},
		{543, false},
		{223450, false},
	}

	var result bool
	var label string

	for _, tt := range tests {
		label = fmt.Sprintf("%d-%v", tt.num, tt.expected)
		t.Run(label, func(t *testing.T) {
			result = digitsNeverDecrease(tt.num)
			if tt.expected {
				assert.True(t, result)
			} else {
				assert.False(t, result)
			}
		})
	}
}
