package main

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	curr := dummy
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}
