package main

import "fmt"

// longestPalindrome returns the longest palindromic substring in the given string
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	// Iterate through each character in the string
	for i := 0; i < len(s); i++ {
		// Odd-length palindrome (center at a single character)
		len1 := expandFromCenter(s, i, i)
		// Even-length palindrome (center between two characters)
		len2 := expandFromCenter(s, i, i+1)

		// Get the maximum length of palindrome found
		maxLen := max(len1, len2)

		// Update start and end indices for the longest palindrome
		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	// Return the substring corresponding to the longest palindrome
	return s[start : end+1]
}

// expandFromCenter checks the length of the palindrome by expanding around the center
func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// Return the length of the palindrome
	return right - left - 1
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println(longestPalindrome("babad")) // Output: "bab" or "aba"
	fmt.Println(longestPalindrome("cbbd"))  // Output: "bb"
	fmt.Println(longestPalindrome("a"))     // Output: "a"
	fmt.Println(longestPalindrome("ac"))    // Output: "a" or "c"
	fmt.Println(longestPalindrome("racecar")) // Output: "racecar"
}
