package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	// Initialize two pointers and a variable to track the maximum area
	left, right := 0, len(height)-1
	maxArea := 0

	// Loop until the pointers meet
	for left < right {
		//  Calculate the current area
		h := int(math.Min(float64(height[left]), float64(height[right])))
		width := right - left
		currentArea := h * width

		// Update the maximum area if the current one is larger
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// Move the pointer pointing to the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	// Test cases
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // Output: 49
	fmt.Println(maxArea([]int{1, 1}))                      // Output: 1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             // Output: 16
	fmt.Println(maxArea([]int{1, 2, 1}))                   // Output: 2
}
