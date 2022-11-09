package 反转系列

type ListNode struct {
	Val  int
	Next *ListNode
}

var successor *listNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	rear := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = rear

	return rear

}
