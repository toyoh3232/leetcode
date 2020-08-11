package p25

import "leetcode/structure"

// ListNode ..
type ListNode = structure.ListNode

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// Test ..
func Test() {
	n7 := &ListNode{Val: 7, Next: nil}
	n6 := &ListNode{Val: 6, Next: n7}
	n5 := &ListNode{Val: 5, Next: n6}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	n1.Print()
	n1 = reverseKGroup(n1, 3)
	n1.Print()
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	var newhead *ListNode
	cursub := head
	//	for cur
	for cursub != nil {
		nextsub := indexOf(cursub, k)
		if cursublast := indexOf(cursub, k-1); cursublast != nil {
			cursublast.Next = nil
			newhead = concat(newhead, reverse(cursub))
		} else {
			newhead = concat(newhead, cursub)
		}
		cursub = nextsub

	}
	return newhead
}

func length(list *ListNode) int {
	if list == nil {
		return 0
	}
	return 1 + length(list.Next)
}

func reverse(list *ListNode) *ListNode {
	var pre, cur, lst *ListNode
	cur = list

	for cur != nil {
		lst = cur.Next
		cur.Next = pre
		pre = cur
		cur = lst
	}
	return pre
}

func concat(list *ListNode, list2 *ListNode) *ListNode {
	if list == nil {
		return list2
	}
	if list2 == nil {
		return list
	}
	head := list
	for list.Next != nil {
		list = list.Next
	}
	list.Next = list2
	return head
}
func indexOf(list *ListNode, index int) *ListNode {
	if list == nil || index == 0 {
		return list
	}
	return indexOf(list.Next, index-1)
}

// @lc code=end
