package p23

import "leetcode/structure"

// ListNode ..
type ListNode = structure.ListNode

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// Test ..
func Test() {

}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeKLists(append(lists, mergeLists(lists[0], lists[1]))[2:])
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val <= list2.Val {
		list1.Next = mergeLists(list1.Next, list2)
		head = list1
	} else if list1.Val > list2.Val {
		list2.Next = mergeLists(list1, list2.Next)
		head = list2
	}
	return head

}

// @lc code=end
