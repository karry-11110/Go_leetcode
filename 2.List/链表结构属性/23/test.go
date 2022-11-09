package main

import (
	"fmt"
)

//单链表结点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

//构造一条单链表，返回链表的第一个结点
func NewList(nums []int) *ListNode {
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
	return head
}
func main() {
	l1 := NewList([]int{1, 1, 2, 3, 4})
	l2 := NewList([]int{0, 5})
	l3 := NewList([]int{1, 4, 7})
	lists := []*ListNode{l1, l2, l3}
	curr := mergeKLists(lists)

	var result []int
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	fmt.Println(result)
}
