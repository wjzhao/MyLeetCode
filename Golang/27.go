package main

import "fmt"

func removeElem(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			fmt.Println(nums)
			size--
			i--
		}
	}
	return size
}

func removeElemV2(nums []int, val int) int {
	p, q := 0, 0
	for q = 0; q < len(nums); q++ {
		if nums[q] != val {
			nums[p] = nums[q]
			p++
		}
	}
	fmt.Println(nums)
	return p
}
