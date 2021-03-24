package p61

import . "leetcode/structure"

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

func Test() int {
	l1 := NewList(0)
	l1.Push(1).Push(2).Push(3)
	l1.Print()
	rotateRight(l1, 2)
	l1.Print()
	return 0
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	return head.Rotate(k)
}

// @lc code=end
