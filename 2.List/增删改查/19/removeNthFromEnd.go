package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	curr := dummy
	prev := dummy
	for curr.Next != nil && n > 0 {
		curr = curr.Next
		n--
	}
	for curr.Next != nil {
		curr = curr.Next
		prev = prev.Next
	}
	prev.Next = prev.Next.Next

	return dummy.Next
}
