package main

//借助哈希表，这里注意判断元素是否在哈希表中的方法
func hasCycle(head *ListNode) bool {
	tmps := map[*ListNode]bool{}
	for head != nil {
		if _, ok := tmps[head]; ok {
			return true
		}
		tmps[head] = true
		head = head.Next
	}
	return false
}

//快慢指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
