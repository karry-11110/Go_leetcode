package main

import "fmt"

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
	l1 := NewList([]int{7, 7, 8, 7, 1})
	curr := middleNode(l1)

	var result []int
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	fmt.Println(result)
}
