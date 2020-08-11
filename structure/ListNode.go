package structure

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// Length ...
func (list *ListNode) Length() int {
	if list == nil {
		return 0
	}
	return 1 + list.Next.Length()
}

// Print ...
func (list *ListNode) Print() {
	if list == nil {
		fmt.Println("nil")
	}
	if list.Next == nil {
		fmt.Println(list.Val)
	} else {
		fmt.Printf("%d->", list.Val)
		list.Next.Print()
	}
}

// Reverse ...
func (list *ListNode) Reverse() *ListNode {
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

// Concat ...
func (list *ListNode) Concat(list2 *ListNode) *ListNode {
	head := list
	for list.Next != nil {
		list = list.Next
	}
	list.Next = list2
	return head
}

// IndexOf ...
func (list *ListNode) IndexOf(index int) *ListNode {
	if index == 0 {
		return list
	}
	return list.Next.IndexOf(index - 1)
}
