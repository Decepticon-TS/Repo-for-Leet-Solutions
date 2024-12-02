func twoSum(nums []int, target int) []int {
    // Create a map to store the difference and its index
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if index, ok := seen[complement]; ok {
            return []int{index, i}
        }
        seen[num] = i
    }

    return nil // This line should never be reached due to the problem constraints
}
