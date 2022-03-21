package main

func mergeTwoList(root1, root2 *LNode) *LNode {
	preRoot := &LNode{}
	cRoot := preRoot
	for root1 != nil && root2 != nil {
		if root1.Val < root2.Val {
			cRoot.Next = root1
			root1 = root1.Next
		} else {
			cRoot.Next = root2
			root2 = root2.Next
		}
		cRoot = cRoot.Next
	}
	if root1 != nil {
		cRoot.Next = root1
	}
	if root2 != nil {
		cRoot.Next = root2
	}
	return preRoot.Next
}
