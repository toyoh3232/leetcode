package p24

import "leetcode/structure"

// ListNode ..
type ListNode = structure.ListNode

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
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
 * }
 */
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	first, second := head, head.Next
	first.Next = second.Next
	second.Next = first
	iter := first

	for iter.Next != nil && iter.Next.Next != nil {
		first = iter.Next
		second = iter.Next.Next
		first.Next = second.Next
		second.Next = first
		iter.Next = second

		iter = iter.Next.Next
	}
	return newHead
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur *ListNode
	newHead := head.Next
	// current
	cur = head
	for cur != nil && cur.Next != nil {
		// do word
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		if pre != nil {
			pre.Next = next
		}
		// next
		pre = cur
		cur = cur.Next
	}
	return newHead
}

// @lc code=end
