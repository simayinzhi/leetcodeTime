package main

import (
	. "computerBase/code"
)

func swapPairs(root *ListNode) *ListNode {
	head := &ListNode{}
	head.Next = root
	node := head
	for node != nil {
		next := node.Next
		if next != nil && next.Next != nil {
			node.Next = next.Next
			next.Next = next.Next.Next
			node.Next.Next = next
			node = next
		} else {
			break
		}
	}
	return head.Next
}

func main() {
	//list := CreateList(1, 2, 3, 4, 5, 6, 10, 9)
	list := CreateList(1)
	PrintList(swapPairs(list))

}
