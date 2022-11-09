package main

import "container/list"

//递归写法
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//终止条件
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	//左
	left := isSameTree(p.Left, q.Left)
	//右
	right := isSameTree(p.Right, q.Right)
	//中
	result := left && right
	return result
}

//迭代
func isSameTree_1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	queue1, queue2 := list.New(), list.New()
	queue1.PushBack(p)
	queue2.PushBack(q)
	for queue1.Len() > 0 && queue2.Len() > 0 {
		node1 := queue1.Remove(queue1.Front()).(*TreeNode)
		node2 := queue2.Remove(queue2.Front()).(*TreeNode)
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue1.PushBack(node1.Left)
		queue1.PushBack(node1.Right)
		queue2.PushBack(node2.Left)
		queue2.PushBack(node2.Right)
	}
	return true
}
