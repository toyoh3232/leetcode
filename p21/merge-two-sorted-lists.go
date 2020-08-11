package p21

import "leetcode/structure"

// ListNode ..
type ListNode = structure.ListNode

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// Test ...
func Test() {

}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	for l1 != nil && l2 != nil {
		var newNode = new(ListNode)
		if l1.Val <= l2.Val {
			newNode.Val = l1.Val
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			newNode.Val = l2.Val
			l2 = l2.Next

		}
		l3 = addNodes(l3, newNode)
	}
	if l1 != nil {
		l3 = addNodes(l3, l1)
	} else if l2 != nil {
		l3 = addNodes(l3, l2)
	}

	return l3

}

func addNodes(list *ListNode, node *ListNode) *ListNode {
	if list == nil {
		return node
	}
	var old = list
	for list.Next != nil {
		list = list.Next
	}
	list.Next = node
	return old
}

// @lc code=end
