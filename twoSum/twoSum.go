package twosum

// O(n²)
func TwoSum1(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// O(n)
func TwoSum2(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := indexMap[complement]; found {
			return []int{index, i}
		}

		indexMap[num] = i
	}

	return nil
}
