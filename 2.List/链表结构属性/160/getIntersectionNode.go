package main

//借助哈希表，但空间复杂度高
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

//快慢指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	lenA, lenB := 0, 0
	//求A,B长度
	for pa != nil {
		pa = pa.Next
		lenA++
	}
	for pb != nil {
		pb = pb.Next
		lenB++
	}

	var step int
	var fast, slow *ListNode
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
