package main

func inorder(root *TreeNode, result *[]int) {
	for root != nil {
		inorder(root.Left, result)
		*result = append(*result, root.Val)
		inorder(root.Right, result)
	}
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorder(root, &result)
	return result
}

func inorderTraversalV2(root *TreeNode) (ret []int) {
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		for root != nil {
			inorder(root.Left)
			ret = append(ret, root.Val)
			inorder(root.Right)
		}
	}
	inorder(root)
	return
}
