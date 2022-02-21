package main

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxResult := 0
	for i := 0; i < len(nums); i++ {
		if i > maxResult {
			return false
		}
		maxResult = max(maxResult, nums[i])
	}
	return true
}
