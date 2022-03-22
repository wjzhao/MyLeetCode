package main

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	backtrackSum3(k, n, 1, []int{}, &result)
	return result
}

func sum(arr []int) int {
	ret := 0
	for i := 0; i < len(arr); i++ {
		ret += arr[i]
	}
	return ret
}

func backtrackSum3(k, n int, start int, path []int, result *[][]int) {
	if len(path) == k && sum(path) == n {
		temp := make([]int, k)
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := start; i <= 9; i++ {
		path = append(path, i)
		backtrackSum3(k, n, i+1, path, result)
		path = path[:len(path)-1]
	}
}
