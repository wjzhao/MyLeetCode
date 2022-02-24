package main

func removeNthFromEnd(root *LNode, target int) *LNode {
	list := []*LNode{root}
	for root.Next != nil {
		list = append(list, root.Next)
		root = root.Next
	}
	index := len(list) - target
	list[index-1].Next = list[index+1]
	return list[0]
}
