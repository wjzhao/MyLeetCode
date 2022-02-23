package main

// [1, 2]
// [[1, 2], [2, 1]]

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	path := []int{}
	result := [][]int{}
	dfs(nums, path, used, &result)
	return result
}

func dfs(nums []int, path []int, used []bool, result *[][]int) {
	if len(path) == len(nums) {
		*result = append(*result, path)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			dfs(nums, path, used, result)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
