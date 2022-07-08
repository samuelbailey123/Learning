package main

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	top := &ListNode{
		Next: head,
	}
	track := top
	for track.Next != nil {
		last, after := reverse(track.Next, k)
		if last == nil {
			return top.Next
		}
		track.Next.Next = after
		track.Next, track = last, track.Next
	}
	return top.Next
}

func reverse(n *ListNode, remaining int) (last, after *ListNode) {
	if remaining == 1 {
		return n, n.Next
	}
	if n.Next == nil {
		return nil, nil
	}
	last, after = reverse(n.Next, remaining-1)
	if last == nil {
		return nil, nil
	}
	n.Next.Next = n
	return
}
