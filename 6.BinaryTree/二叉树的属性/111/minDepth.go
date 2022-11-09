package main

import (
	"container/list"
)

//迭代：层次遍历
func minDepth(root *TreeNode) int {
	result := 0
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return result + 1
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result++

	}
	return result
}

//递归：
func minDepth_1(root *TreeNode) int {
	//终止条件
	if root == nil {
		return 0
	}
	//左
	left := minDepth(root.Left)
	//右
	right := minDepth(root.Right)
	//中
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	result := min(left, right) + 1
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
