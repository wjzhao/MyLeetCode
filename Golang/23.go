package main

func mergeKLists(lists []*LNode) *LNode {
	middle := len(lists) / 2
	left := mergeKLists(lists[:middle])
	right := mergeKLists(lists[middle:])
	return mergeTwoList(left, right)
}
