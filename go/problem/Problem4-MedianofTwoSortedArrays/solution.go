package main

import "fmt"

// findMedianSortedArrays finds the median of two sorted arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array for optimized binary search
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)

	left, right := 0, m

	for left <= right {
		// Partition indices for nums1 and nums2
		i := (left + right) / 2
		j := (m+n+1)/2 - i

		// Handle edge cases for boundaries
		nums1LeftMax := -1_000_001
		if i > 0 {
			nums1LeftMax = nums1[i-1]
		}

		nums1RightMin := 1_000_001
		if i < m {
			nums1RightMin = nums1[i]
		}

		nums2LeftMax := -1_000_001
		if j > 0 {
			nums2LeftMax = nums2[j-1]
		}

		nums2RightMin := 1_000_001
		if j < n {
			nums2RightMin = nums2[j]
		}

		// Check if valid partition
		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			// Calculate median based on combined length
			if (m+n)%2 == 0 {
				return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
			} else {
				return float64(max(nums1LeftMax, nums2LeftMax))
			}
		} else if nums1LeftMax > nums2RightMin {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	// No valid partition found (shouldn't reach here)
	return -1.0
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // Output: 2.00000
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // Output: 2.50000
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0})) // Output: 0.00000
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))        // Output: 1.00000
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))        // Output: 2.00000
}
