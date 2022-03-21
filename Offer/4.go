package main

func singleNumber(nums []int) int {
	temp := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				continue
			}
		}
		temp = i
	}
	return nums[temp]
}
