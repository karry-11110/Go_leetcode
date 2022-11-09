package main

//借助切片加双指针
func isPalindrome(head *ListNode) bool {
	tmps := []int{}
	for ; head != nil; head = head.Next {
		tmps = append(tmps, head.Val)
	}
	n := len(tmps)
	low, high := 0, n-1
	for low < high {
		if tmps[low] != tmps[high] {
			return false
		}
		low++
		high--
	}
	return true
}

func findFristHalfEnd(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fristHalfEnd := findFristHalfEnd(head)
	secondHalfStart := reverseList(fristHalfEnd.Next)
	p1, p2 := head, secondHalfStart
	result := true
	for result && p2 != nil { //注意这里为什么要p2不为空而不是p1
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	fristHalfEnd.Next = reverseList(secondHalfStart)
	return result

}
