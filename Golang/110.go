package main

func deepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(deepth(root.Left), deepth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := deepth(root.Left)
	rightHeight := deepth(root.Right)
	return abs(leftHeight-rightHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
