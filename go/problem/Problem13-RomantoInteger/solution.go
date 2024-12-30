package main

import "fmt"

func romanToInt(s string) int {
	// Map for Roman numerals including subtractive cases
	values := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	n := len(s)
	total := values[s[n-1]] // Start with the last numeral's value

	// Iterate from second last to the first numeral
	for i := n - 2; i >= 0; i-- {
		// If current value is less than the next value, subtract it
		if values[s[i]] < values[s[i+1]] {
			total -= values[s[i]]
		} else {
			total += values[s[i]]
		}
	}

	return total
}

func main() {
	// Test cases
	fmt.Println(romanToInt("III"))     // Output: 3
	fmt.Println(romanToInt("LVIII"))   // Output: 58
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
	fmt.Println(romanToInt("IV"))      // Output: 4
	fmt.Println(romanToInt("XL"))      // Output: 40
	fmt.Println(romanToInt("CDXLIV"))  // Output: 444
}
