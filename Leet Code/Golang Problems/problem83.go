package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var curr, prev *ListNode
	curr = head
	prev = head
	for curr != nil {
		if curr.Val == prev.Val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return head
}

func main() {
	var head *ListNode
	head = &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	fmt.Println(deleteDuplicates(head))
}
