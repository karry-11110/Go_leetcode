package main

//两升序合并成一个升序
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

//两升序合并成一个降序
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}

	p1, p2, p := l1, l2, dummy
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			tmp2 := p2
			p2 = tmp2.Next
			tmp2.Next = p.Next
			p.Next = tmp2
		} else {
			tmp1 := p1
			p1 = tmp1.Next
			tmp1.Next = p.Next
			p.Next = tmp1

		}

	}
	if p1 != nil {
		for p1 != nil {
			tmp1 := p1
			p1 = tmp1.Next
			tmp1.Next = p.Next
			p.Next = tmp1
		}

	}
	if p2 != nil {
		for p2 != nil {
			tmp2 := p2
			p2 = tmp2.Next
			tmp2.Next = p.Next
			p.Next = tmp2
		}
	}
	return dummy.Next
}
