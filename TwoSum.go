package main

func twoSum(nums []int, target int) []int {

	var b []int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && (nums[i] + nums[j]) == target{
				b = []int{i, j}
				return b
			}
		}
	}

	return b
}