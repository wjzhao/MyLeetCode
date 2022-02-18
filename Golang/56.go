package main

// [1, 2]
// [0, 5] [3, 5] [4, 9]
// {{1, 3}, {2, 6}, {8, 10}, {15, 18}}
// [[1, 3], ]

func mergeList(intervals [][]int) [][]int {
	ret := [][]int{intervals[0]}
	current := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > ret[current][1] {
			ret = append(ret, intervals[i])
			current++
		} else {
			ret[current][0] = intervals[i][0]
			ret[current][1] = max(ret[current][1], intervals[i][1])
		}
	}
	return ret
}
