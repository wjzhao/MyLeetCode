package main

// nums升序排列
func search1(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			l, r := mid, mid
			for l > 0 && nums[l-1] == target {
				l--
			}
			for r < len(nums)-1 && nums[r+1] == target {
				r++
			}
			return []int{l, r}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
