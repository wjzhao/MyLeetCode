package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	start, end, index, sum, current, diff := 0, 0, 0, 0, 0, 999999
	for index = 1; index < len(nums)-1; index++ {
		start, end = 0, len(nums)-1
		for start < index && end > index {
			sum = nums[start] + nums[index] + nums[end]
			if abs(sum-target) < diff {
				current, diff = sum, abs(sum-target)
			}
			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return current
}
