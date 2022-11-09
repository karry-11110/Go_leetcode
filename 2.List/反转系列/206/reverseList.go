package main

//双指针法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode //默认是nil了
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

//前序递归：可以和双指针互改
func reverseList(head *ListNode) *ListNode {
	return helper(nil, head)
}
func helper(pre *ListNode, curr *ListNode) *ListNode {
	if curr == nil {
		return pre
	}
	next := curr.Next
	curr.Next = pre
	return helper(curr, next)
}

// //递归法:后序递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rear := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rear
}
