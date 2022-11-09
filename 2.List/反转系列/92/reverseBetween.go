package main

//双指针法翻转前n个节点
func reverseN(head *ListNode, n int) *ListNode {
	var pre *ListNode
	curr := head
	if n == 0 {
		return head
	}
	for curr != nil && n >= 1 {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
		n--
	}
	successor := curr
	head.Next = successor
	return pre
}

// //递归法翻转前n个节点
var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	rear := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return rear
}

// //递归法翻转区域链表:需要借助反转前n个节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	rear := reverseBetween(head.Next, left-1, right-1)
	head.Next = rear
	return head
}

//双指针反转区域链表：借助反转前n个节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	rear := reverseN(pre.Next, right-left+1)
	pre.Next = rear
	return dummy.Next
}

//头插法:pre这个过程中不变
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	curr := pre.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next

}
