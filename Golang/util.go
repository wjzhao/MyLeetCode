package main

type LNode struct {
	Val  int
	Next *LNode
}

type DNode struct {
	Val       int
	Pre, Next *DNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

func printList(head *LNode) (ret []int) {
	for {
		ret = append(ret, head.Val)
		if head.Next == nil {
			return
		} else {
			head = head.Next
		}
	}
}

func makeList(list []int) *LNode {
	head := LNode{Val: list[0]}
	current := &head
	for i := 1; i < len(list); i++ {
		next := LNode{Val: list[i]}
		current.Next = &next
		current = current.Next
	}
	return &head
}
