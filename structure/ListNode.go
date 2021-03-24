package structure

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewList ...
func NewList(initial int) *ListNode {
	first := &ListNode{Val: initial}
	return first
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
	} else if list.Next == nil {
		fmt.Println(list.Val)
	} else {
		fmt.Printf("%d->", list.Val)
		list.Next.Print()
	}
}

// Reverse ...
func (list *ListNode) Reverse() *ListNode {
	first := list
	cur := list
	if list.Next == nil {
		return list
	}
	snd := list.Next
	if snd.Next == nil {
		list.Val, snd.Val = snd.Val, list.Val
		return first
	}
	var pre, next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	last := pre
	first.Next, last.Next = last.Next, first.Next
	first.Val, last.Val = last.Val, first.Val
	snd.Next = last
	return first

}

// Concat ...
func (list *ListNode) Concat(list2 *ListNode) *ListNode {
	first := list
	for list.Next != nil {
		list = list.Next
	}
	list.Next = list2
	return first
}

// IndexOf ...
func (list *ListNode) IndexOf(index int) int {
	if index == 0 {
		return list.Val
	}
	return list.Next.IndexOf(index - 1)
}

// Push ...
func (list *ListNode) Push(v int) *ListNode {
	first := list
	next := new(ListNode)
	next.Val = v
	next.Next = nil
	list.Concat(next)
	return first
}

// Rotate ...
func (list *ListNode) Rotate(k int) *ListNode {
	isR := false
	first := list
	if k > 0 {
		isR = true
	} else {
		k = -k
	}
	if isR {
		list.Reverse()
	}
	k = k % list.Length()
	for k > 0 {
		val := list.Val
		cur := list
		for cur.Next != nil {
			cur.Val = cur.Next.Val
			cur = cur.Next
		}
		cur.Val = val
		k--
	}
	if isR {
		list.Reverse()
	}
	return first
}
