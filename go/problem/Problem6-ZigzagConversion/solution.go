package main

import (
	"fmt"
)

// convert rearranges the input string `s` into the zigzag pattern with `numRows` rows
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Initialize the result slice with the exact required size
	n := len(s)
	result := make([]byte, 0, n)

	// Calculate the step for the zigzag pattern
	step := 2*numRows - 2

	// Build the result row by row
	for row := 0; row < numRows; row++ {
		for i := row; i < n; i += step {
			// Append the character in the current row
			result = append(result, s[i])
			// Handle the diagonal characters for middle rows
			if row > 0 && row < numRows-1 {
				diag := i + step - 2*row
				if diag < n {
					result = append(result, s[diag])
				}
			}
		}
	}

	return string(result)
}

func main() {
	// Test cases
	fmt.Println(convert("PAYPALISHIRING", 3)) // Output: "PAHNAPLSIIGYIR"
	fmt.Println(convert("PAYPALISHIRING", 4)) // Output: "PINALSIGYAHRPI"
	fmt.Println(convert("A", 1))              // Output: "A"
	fmt.Println(convert("ABCDE", 2))          // Output: "ACEBD"
}
