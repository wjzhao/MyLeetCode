package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	start, end, index, sum, result := 0, 0, 0, 0, [][]int{}
	for index = 1; index < len(nums)-1; index++ {
		start, end = 0, len(nums)-1
		for start < index && end > index {
			sum = nums[start] + nums[index] + nums[end]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func threeSumV2(nums []int) [][]int {
	start, end, index, sum, result := 0, 0, 0, 0, [][]int{}
	for index = 1; index < len(nums)-1; index++ {
		start, end = 0, len(nums)-1
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end++
				continue
			}
			sum = nums[start] + nums[index] + nums[end]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
