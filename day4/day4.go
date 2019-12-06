package main

import "fmt"

const (
	start = 136818
	end   = 685979
)

type Rule func(int) bool

// compareDigits compares digits in a number one-by-one, applying the condition function
// and returning with successResult if the condition is true for any two digits.
// returns the negation of successResult at the end if condition never was true
func compareDigits(num int, condition func(int, int) bool, successResult bool, lastCondition func() bool) bool {
	var remain int
	prevRemain := num % 10
	for num > 1 {
		num /= 10
		remain = num % 10
		if condition(prevRemain, remain) {
			return successResult
		}
		prevRemain = remain
	}
	return lastCondition()
}

// Two adjacent digits are the same (like 22 in 122345).
func twoAdjacentDigitsAreTheSameAtleastOnce(num int) bool {
	count := 1

	conditionFunc := func(first, second int) bool {
		if count == 2 && first != second {
			return true
		} else if first != second {
			count = 1
		} else if first == second {
			count++
		}
		return false
	}

	// we need to check the last two digits, because there will be no more
	// but we can't know that inside the conditionFunc
	lastConditionFunc := func() bool { return count == 2 }

	return compareDigits(num, conditionFunc, true, lastConditionFunc)
}

// Going from left to right, the digits never decrease
// they only ever increase or stay the same (like 111123 or 135679)
func digitsNeverDecrease(num int) bool {
	conditionFunc := func(first, second int) bool { return first < second }
	lastConditionFunc := func() bool { return true }
	return compareDigits(num, conditionFunc, false, lastConditionFunc)
}

func allRulesApply(password int, rules []Rule) bool {
	for _, rule := range rules {
		if !rule(password) {
			return false
		}
	}
	return true
}

func countPasswords(start, end int, rules []Rule) int {
	count := 0
	for password := start; password <= end; password++ {
		if allRulesApply(password, rules) {
			count++
		}
	}
	return count
}

func main() {
	var rules = []Rule{
		// "six-digit" and "value within the given range" rules
		// are implicit by using start and end input
		twoAdjacentDigitsAreTheSameAtleastOnce,
		digitsNeverDecrease,
	}
	res := countPasswords(start, end, rules)
	fmt.Println("Result:", res)
}
