package main

import "fmt"

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	// Map to store the last seen index of each character
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		// If the character is already in the map and is within the current window
		if index, exists := charIndex[s[end]]; exists && index >= start {
			// Move the start pointer to the right of the last occurrence of the current character
			start = index + 1
		}
		// Update the character's last seen index
		charIndex[s[end]] = end
		// Update the max length of the substring
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
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
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
	fmt.Println(lengthOfLongestSubstring(" "))        // Output: 1
	fmt.Println(lengthOfLongestSubstring("au"))       // Output: 2
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // Output: 3
}
