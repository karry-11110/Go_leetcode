package main

//借助哈希表
func lengthCycle(head *ListNode) int {
	tmps := map[*ListNode]bool{}
	for head != nil {
		if _, ok := tmps[head]; ok {
			p1, p2 := head, head
			length := 1
			for p2.Next != p1 {
				p2 = p2.Next
				length++
			}
			return length
		}
		tmps[head] = true
		head = head.Next
	}
	return 0
}

//快慢指针
func lengthCycle(head *ListNode) int {
	slow, fast := head, head
	slowStep, fastStep := 0, 0
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		slowStep++
		fast = fast.Next.Next
		fastStep = fastStep + 2
		if fast == slow {
			return fastStep - slowStep
		}
	}
	return 0
}
