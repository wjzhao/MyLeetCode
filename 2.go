package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNum(l1, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			n1 = 0
			// carry = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
			// carry = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

var n6 = ListNode{2, &n5}
var n5 = ListNode{4, &n4}
var n4 = ListNode{3, nil}
var n3 = ListNode{5, &n2}
var n2 = ListNode{6, &n1}
var n1 = ListNode{4, nil}

func addTwoNumTest() {
	node := addTwoNum(&n6, &n3)
	for node.Next != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)
}
