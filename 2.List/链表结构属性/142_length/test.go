package main

import "fmt"

//单链表结点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

//构造环形单链表
func NewList(nums []int, pos int) *ListNode {
	//先构造一条单链表
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	//构造环
	if pos >= 0 {
		p := head
		for pos > 0 {
			p = p.Next
			pos--
		}
		cur.Next = p
	}

	return head
}
func main() {
	l1 := NewList([]int{1}, 0)
	curr := lengthCycle(l1)
	result := curr
	fmt.Println(result)
}
