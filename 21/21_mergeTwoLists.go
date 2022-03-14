package main

import (
	. "computerBase/code"
	"math"
)

//if list1[0]<list2[0]
//list1[0]+merge(list1[1:],list2)
//else
//list2[0]+merge(list1,list2[1:])
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else {
		if list1.Val < list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists(list2.Next, list1)
			return list2
		}
	}
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	node := result
	for list1 != nil || list2 != nil {
		var var1 = math.MaxInt64
		var var2 = math.MaxInt64
		if list1 != nil {
			var1 = list1.Val
		}
		if list2 != nil {
			var2 = list2.Val
		}
		if var1 > var2 {
			node.Next = &ListNode{Val: var2}
			list2 = list2.Next
		} else {
			node.Next = &ListNode{Val: var1}
			list1 = list1.Next
		}
		node = node.Next
	}
	return result.Next
}

func main() {
	var list1 = CreateList(1, 2, 3, 4, 5, 8)
	var list2 = CreateList(1, 5, 7, 9, 10, 18)
	var list3 = mergeTwoLists(list1, list2)
	PrintList(list3)
}
