package main

import "container/list"

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//左
	leftSum := sumOfLeftLeaves(root.Left)
	//右
	rightSum := sumOfLeftLeaves(root.Right)
	//中
	midValue := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		midValue = root.Left.Val
	}
	sum := leftSum + rightSum + midValue
	return sum
}

func sumOfLeftLeaves_1(root *TreeNode) int {
	result := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return result
}
