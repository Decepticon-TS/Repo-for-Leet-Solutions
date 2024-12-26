package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	index := 0
	sign := 1
	result := 0

	// Step 1: Skip leading whitespace
	for index < n && unicode.IsSpace(rune(s[index])) {
		index++
	}

	// Step 2: Handle optional '+' or '-'
	if index < n && (s[index] == '+' || s[index] == '-') {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}

	// Step 3: Convert characters to integer
	for index < n && unicode.IsDigit(rune(s[index])) {
		digit := int(s[index] - '0')

		// Step 4: Check for overflow/underflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		index++
	}

	return result * sign
}

func main() {
	// Test cases
	fmt.Println(myAtoi("42"))           // Output: 42
	fmt.Println(myAtoi("   -42"))       // Output: -42
	fmt.Println(myAtoi("4193 with words")) // Output: 4193
	fmt.Println(myAtoi("words and 987"))   // Output: 0
	fmt.Println(myAtoi("-91283472332"))    // Output: -2147483648 (Clamped to math.MinInt32)
}
