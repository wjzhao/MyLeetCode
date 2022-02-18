package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
