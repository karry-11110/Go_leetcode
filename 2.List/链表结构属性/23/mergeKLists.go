package main

func mergeKLists(lists []*ListNode) *ListNode {
	listsLength := len(lists)
	var pre *ListNode
	for i := 0; i < listsLength; i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		curr := lists[i]
		pre = mergeTwoLists(pre, curr)

	}
	return pre
}
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	p1, p2, p := l1, l2, dummy
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
