package main

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preorder(root, res)
	return res
}

func preorder(tree *TreeNode, result []int) {
	if tree != nil {
		result = append(result, tree.Val)
		preorder(tree.Left, result)
		preorder(tree.Right, result)
	}
}
