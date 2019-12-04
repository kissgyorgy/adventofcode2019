package main

import "fmt"

const (
	start = 136818
	end   = 685979
)

type Rule func(int) bool

func twoAdjacentDigitsAreTheSame(num int) bool {
	return false
}

func digitsNeverDecrease(num int) bool {
	var remain int
	prevRemain := num % 10
	for num > 1 {
		num /= 10
		remain = num % 10
		if prevRemain < remain {
			return false
		}
		prevRemain = remain
	}
	return true
}

func countPasswords(start, end int, rules []Rule) int {
	count := 0
	for password := start; password <= end; password++ {
		for _, rule := range rules {
			if rule(password) {
				count++
			}
		}
	}
	return count
}

func main() {
	var rules = []Rule{
		// "six-digit" and "value within the given range" rules
		// are implicit by using start and end input
		twoAdjacentDigitsAreTheSame,
		digitsNeverDecrease,
	}
	res := countPasswords(start, end, rules)
	fmt.Println("Result:", res)
}
