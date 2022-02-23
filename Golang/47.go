package main

func permuteUnique(nums []int) [][]int {
	used := make([]bool, len(nums))
	result := [][]int{}
	path := []int{}
	dfsV2(nums, path, used, &result)
	return result
}

// 带有剪枝操作的DFS
func dfsV2(nums, path []int, used []bool, result *[][]int) {
	if len(path) == len(nums) {
		*result = append(*result, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfsV2(nums, path, used, result)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
