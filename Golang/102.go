package main

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		temp := []int{}
		tempQueue := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			temp = append(temp, queue[j].Val)
			if queue[j].Left != nil {
				tempQueue = append(tempQueue, queue[j].Left)
			}
			if queue[j].Right != nil {
				tempQueue = append(tempQueue, queue[j].Right)
			}
		}
		result = append(result, temp)
		queue = tempQueue
	}
	return result
}
