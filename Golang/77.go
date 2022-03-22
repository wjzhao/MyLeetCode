package main

func combine(n int, k int) [][]int {
	result := [][]int{}
	backtrack(n, k, 1, []int{}, &result)
	return result
}

func backtrack(n, k, start int, path []int, result *[][]int) {
	if len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		backtrack(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}
