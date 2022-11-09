package main

import "container/list"

//迭代：层次遍历
func maxDepth(root *TreeNode) int {
	result := [][]int{}
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		subresult := []int{}
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			subresult = append(subresult, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, subresult)
	}
	return len(result)
}

//递归：后序遍历
func maxDepth_1(root *TreeNode) int {
	//终止条件
	if root == nil {
		return 0
	}
	//左
	maxDepthLeft := maxDepth(root.Left)
	//右
	maxDepthRight := maxDepth(root.Right)
	//中
	result := max(maxDepthLeft, maxDepthRight) + 1
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
