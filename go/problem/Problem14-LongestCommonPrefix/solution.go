package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the initial prefix
	prefix := strs[0]

	// Iterate over the rest of the strings
	for _, str := range strs[1:] {
		// If any string is empty, the common prefix must be ""
		if str == "" {
			return ""
		}

		// Shrink the prefix until it matches the start of the current string
		for len(prefix) > 0 && !startsWith(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		// If prefix becomes empty, return ""
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// Helper function to check if str starts with prefix
func startsWith(str, prefix string) bool {
	if len(prefix) > len(str) {
		return false
	}
	return str[:len(prefix)] == prefix
}

func main() {
	// Test cases
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // Output: ""
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))                  // Output: "a"
	fmt.Println(longestCommonPrefix([]string{"abab", "aba", ""}))          // Output: ""
}
