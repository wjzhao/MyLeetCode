package main

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i, j := 0, len(nums)-1
	for p := len(nums) - 1; p >= 0; p-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[p] = nums[i] * nums[i]
			i++
		} else {
			result[p] = nums[j] * nums[j]
			j--
		}
	}
	return result
}
