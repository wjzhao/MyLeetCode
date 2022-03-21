package main

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	current := head
	for current != nil {
		temp := current.Next
		current.Next = pre
		pre = current
		current = temp
	}
	return pre
}
