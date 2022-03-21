package main

import "fmt"

func removeDuplicates(nums []int) int {
	p := 1
	for q := 1; q < len(nums); q++ {
		if nums[q-1] != nums[q] {
			nums[p] = nums[q]
			p++
		}
	}
	fmt.Println(nums)
	return p
}
