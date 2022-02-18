package main

func levelOrderBottom(root *TreeNode) [][]int {
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
	for i, length := 0, len(result); i < length/2; i++ {
		result[i], result[length-i-1] = result[length-i-1], result[i]
	}
	return result
}
