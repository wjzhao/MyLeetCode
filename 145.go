package main

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(tree *TreeNode, result *[]int) {
	if tree != nil {
		postorder(tree.Left, result)
		postorder(tree.Right, result)
		*result = append(*result, tree.Val)
	}
}
