package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	num      int
	expected bool
}

func runRuleTests(t *testing.T, tests []TestCase, ruleFunc Rule) {
	t.Helper()

	var result bool
	var label string

	for _, tt := range tests {
		label = fmt.Sprintf("%d-%v", tt.num, tt.expected)
		t.Run(label, func(t *testing.T) {
			result = ruleFunc(tt.num)
			if tt.expected {
				assert.True(t, result)
			} else {
				assert.False(t, result)
			}
		})
	}
}

func TestDigitsNeverDecrease(t *testing.T) {
	tests := []TestCase{
		{0, true},
		{100, false},
		{102, false},
		{111, true},
		{112, true},
		{123, true},
		{543, false},
		{223450, false},
	}
	runRuleTests(t, tests, digitsNeverDecrease)
}

func TestTwoAdjacentDigitsAreTheSame(t *testing.T) {
	tests := []TestCase{
		{100, true},
		{101, false},
		{12345, false},
		{111111, false},
		{112345, true},
		{123345, true},
		{123789, false},
		{112233, true},
		{123444, false},
		{111122, true},
	}
	runRuleTests(t, tests, twoAdjacentDigitsAreTheSameAtleastOnce)
}
