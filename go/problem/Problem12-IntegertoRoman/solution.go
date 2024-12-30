package main

import "fmt"

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	// Iterate through the values and symbols
	for i := 0; i < len(vals); i++ {
		// While num is greater than or equal to the current value
		for num >= vals[i] {
			result += symbols[i] // Append the corresponding Roman symbol
			num -= vals[i]       // Subtract the value from num
		}
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(intToRoman(3749)) // Output: "MMMDCCXLIX"
	fmt.Println(intToRoman(58))   // Output: "LVIII"
	fmt.Println(intToRoman(1994)) // Output: "MCMXCIV"
	fmt.Println(intToRoman(1))    // Output: "I"
	fmt.Println(intToRoman(3999)) // Output: "MMMCMXCIX"
}
