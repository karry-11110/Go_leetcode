package main

//穿针引线

func reverseK(start *ListNode, k int) *ListNode {
	var pre *ListNode
	curr := start
	if k == 0 {
		return start
	}
	for curr != nil && k >= 1 {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
		k--
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	start, end := head, head
	num := k
	for k > 0 {
		if end == nil {
			return head
		}
		end = end.Next
		k--
	}
	newHead := reverseK(start, num)
	start.Next = reverseKGroup(end, num)
	return newHead
}

//穿针引线
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy
	end := dummy.Next

	for end != nil {
		for i := 1; i < k; i++ { //这里注意end的位置只能在哪？？？
			if end == nil {
				break
			}
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		endNext := end.Next
		pre.Next = reverseK(start, k)
		pre = start
		pre.Next = endNext
		end = endNext
	}
	return dummy.Next
}
