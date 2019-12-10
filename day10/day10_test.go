package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	p1 = Point{1, 0}
	p2 = Point{2, 2}
	p3 = Point{3, 2}
	p4 = Point{3, 4}
	p5 = Point{4, 0}
	p6 = Point{4, 2}
	p7 = Point{4, 3}
)

func TestAreOnTheSameLine(t *testing.T) {
	tests := []struct {
		p1, p2, p3 Point
	}{
		{p1, p4, p2},
		{p7, p6, p5},
		{p6, p7, p5},
		{p5, p7, p6},
		{p1, p3, p7},
	}
	for _, tt := range tests {
		desc := fmt.Sprintf("--%s--%s--%s--", tt.p1, tt.p2, tt.p3)
		t.Run(desc, func(t *testing.T) {
			assert.True(t, areOnTheSameLine(tt.p1, tt.p2, tt.p3))
		})
	}
}

func TestIsBetweenTwoPoints(t *testing.T) {
	tests := []struct {
		first, middle, last Point
	}{
		{p1, p2, p4},
		{p4, p2, p1},
		{p1, p3, p7},
		{p7, p3, p1},
		{p5, p6, p7},
		{p7, p6, p5},
	}
	for _, tt := range tests {
		desc := fmt.Sprintf("%s-->%s-->%s", tt.first, tt.middle, tt.last)
		t.Run(desc, func(t *testing.T) {
			assert.True(t, tt.middle.isBetweenTwoPoints(tt.first, tt.last))
		})
	}
}
