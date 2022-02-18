# 递归求解


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if not l1: return l2
        if not l2: return l1
        head = ListNode(None)
        if l1.val <= l2.val:
            head.val = l1.val
            head.next = self.mergeTwoLists(l1.next, l2)
        else:
            head.val = l2.val
            head.next = self.mergeTwoLists(l1, l2.next)
        return head


def print_node(node: ListNode) -> int:
    print(node.val)
    return print_node(node.next)
