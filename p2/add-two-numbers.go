package p2

import . "leetcode/structure"

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Length() > l2.Length() {
		l1, l2 = l2, l1
	}
	l2o := l2
	carrier := 0
	var pre *ListNode
	for l1 != nil {
		val := l1.Val + l2.Val + carrier
		if val > 9 {
			carrier = 1
		} else {
			carrier = 0
		}
		l1.Val = val % 10
		l2.Val = val % 10
		l1 = l1.Next
		pre = l2
		l2 = l2.Next
	}

	for l2 != nil {
		val := l2.Val + carrier
		if val > 9 {
			carrier = 1
		} else {
			carrier = 0
		}
		l2.Val = val % 10
		pre = l2
		l2 = l2.Next
	}
	if carrier != 0 {
		pre.Next = &ListNode{Val: carrier, Next: nil}
	}
	return l2o
}

// @lc code=end
