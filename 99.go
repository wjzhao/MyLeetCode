package main

func traverse(root *TreeNode, result *([]*TreeNode)) {
	for root != nil {
		traverse(root.Left, result)
		*result = append(*result, root)
		traverse(root.Right, result)
	}
}

func recoverTree(root *TreeNode) {
	result := make([]*TreeNode, 0)
	traverse(root, &result)
	pre := result[0]
	var x, y *TreeNode
	for i := 1; i < len(result); i++ {
		if pre.Val > result[i].Val {
			x = pre
			y = result[i]
		}
		pre = result[i]
	}
	if x != nil && y != nil {
		x, y = y, x
	}
}
