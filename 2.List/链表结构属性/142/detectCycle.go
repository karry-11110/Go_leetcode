package main

//借助哈希表
func detectCycle(head *ListNode) *ListNode {
	tmps := map[*ListNode]bool{}
	for head != nil {
		if _, ok := tmps[head]; ok {
			return head
		}
		tmps[head] = true
		head = head.Next
	}
	return nil
}

//快慢指针
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
