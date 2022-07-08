package main

// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next.Next
		cur.Next.Next = tmp.Next
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = cur.Next.Next
	}
	return dummy.Next
}
