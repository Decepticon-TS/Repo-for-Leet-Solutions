package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	if n < 3 {
		return result
	}

	// Sort the array
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// Skip duplicate elements for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use two-pointer technique
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate elements for the second and third numbers
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move pointers inward
				left++
				right--
			} else if sum < 0 {
				// Increase the sum by moving the left pointer
				left++
			} else {
				// Decrease the sum by moving the right pointer
				right--
			}
		}
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // Output: [[-1, -1, 2], [-1, 0, 1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // Output: []
	fmt.Println(threeSum([]int{0, 0, 0}))             // Output: [[0, 0, 0]]
}
