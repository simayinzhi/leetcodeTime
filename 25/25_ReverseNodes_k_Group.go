package main

import (
	. "computerBase/code"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	mark := head

	for {
		if mark == nil {
			return head
		}
		splitFirst := mark
		for i := 0; i < k-1; i++ {
			if mark.Next == nil {
				return head
			}
			mark = mark.Next
		}
		node := mark
		for i := 0; i < k/2; i++ {
			temp := splitFirst.Val
			splitFirst.Val = node.Val
			node.Val = temp

			splitFirst = splitFirst.Next
			node = splitFirst
			for j := 0; j < k-3-i*2; j++ {
				node = node.Next
			}
		}
		mark = mark.Next
	}
	return head
}

func main() {
	//node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//reverseKGroup(node, 5)
	//PrintList(node)
	node1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	reverseKGroup(node1, 6)
	PrintList(node1)
}
