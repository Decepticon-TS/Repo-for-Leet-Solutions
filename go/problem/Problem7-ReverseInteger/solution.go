package main

import (
	"fmt"
)

func reverse(x int) int {
	const intMin, intMax = -1 << 31, 1<<31 - 1

	result := 0
	for x != 0 {
		// Extract the last digit
		digit := x % 10
		x /= 10

		// Check for overflow/underflow
		if result > (intMax-digit)/10 || result < (intMin-digit)/10 {
			return 0
		}

		// Append the digit
		result = result*10 + digit
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(reverse(123))        // Output: 321
	fmt.Println(reverse(-123))       // Output: -321
	fmt.Println(reverse(120))        // Output: 21
	fmt.Println(reverse(0))          // Output: 0
	fmt.Println(reverse(1534236469)) // Output: 0 (overflow case)
}
