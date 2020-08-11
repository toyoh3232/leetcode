package p19

import "leetcode/structure"

// ListNode ..
type ListNode = structure.ListNode

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// Run ...
func Run() {

}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lt := length(head)
	index := lt - n
	iter := head
	var pre *ListNode
	for i := 0; iter != nil; i, pre, iter = i+1, iter, iter.Next {
		if i == index && i == 0 {
			return iter.Next
		}
		if i == index {
			pre.Next = iter.Next
		}
	}
	return head
}

func length(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + length(head.Next)
}

// @lc code=end
