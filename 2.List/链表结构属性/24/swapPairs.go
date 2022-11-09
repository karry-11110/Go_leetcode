package main

func swapPairs(head *ListNode) *ListNode {
	var dummy = &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy
	curr := head
	for curr != nil && curr.Next != nil {
		next := curr.Next
		curr.Next = next.Next
		next.Next = curr
		pre.Next = next
		pre = curr
		curr = pre.Next
	}
	return dummy.Next
}

//递归实现
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead

}
