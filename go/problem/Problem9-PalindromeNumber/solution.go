package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// Negative numbers and numbers ending in 0 (except 0 itself) cannot be palindromes
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversed := 0
	original := x

	// Reverse the digits of the number
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// Compare the original number with the reversed number
	return original == reversed
}

func main() {
	// Test cases
	fmt.Println(isPalindrome(121))  // Output: true
	fmt.Println(isPalindrome(-121)) // Output: false
	fmt.Println(isPalindrome(10))   // Output: false
	fmt.Println(isPalindrome(0))    // Output: true
}
