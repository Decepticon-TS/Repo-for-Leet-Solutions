package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// Sort the nums array
	sort.Ints(nums)
	n := len(nums)
	closestSum := math.MaxInt32

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// Check if the current sum is closer to the target than the previous closest sum
			if abs(target-sum) < abs(target-closestSum) {
				closestSum = sum
			}

			// Move the pointers based on the comparison between sum and target
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				// If sum == target, it's the closest possible, so return it
				return sum
			}
		}
	}
	return closestSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // Output: 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))      // Output: 0
}
