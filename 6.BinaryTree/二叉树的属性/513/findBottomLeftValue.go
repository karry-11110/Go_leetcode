package main

import "container/list"

//递归
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	maxDeep := 0
	value := 0
	var find func(node *TreeNode, deep int)
	find = func(node *TreeNode, deep int) {
		//终止条件，找到叶子节点,最左边的值在第一个
		if node.Left == nil && node.Right == nil {
			if deep > maxDeep {
				maxDeep = deep
				value = node.Val
			}

		}
		if node.Left != nil {
			find(node.Left, deep+1) //间接回溯
		}
		if node.Right != nil {
			find(node.Right, deep+1)
		}

	}
	find(root, maxDeep)
	return value
}

// 显示递归
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	maxDepth := 0
	depth := 0
	result := 0
	var find func(node *TreeNode)
	find = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				result = node.Val
			}
		}
		depth++
		find(node.Left)
		find(node.Right)
		depth--
	}
	find(root)
	return result
}

//层次遍历
func findBottomLeftValue_1(root *TreeNode) int {
	queue := list.New()
	var gradation int
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				gradation = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return gradation
}
