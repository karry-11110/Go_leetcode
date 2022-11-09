package main

import (
	"container/list"
)

// 解题思路：本题遍历只能是“后序遍历”，因为我们要通过递归函数的返回值来判断两个子树的内侧节点和外侧节点是否相等。
//此题解法并不是严格的树的序列递归，迭代也不是严格的层次遍历
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}
func compare(left, right *TreeNode) bool {
	//终止条件，也就是排除空结点和数值不同的情况
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	//左
	outside := compare(left.Left, right.Right)
	//右
	inside := compare(left.Right, right.Left)
	//中
	isSame := outside && inside
	return isSame
}

//迭代解法:类似于层次遍历
func isSymmetric_1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		leftNode := queue.Remove(queue.Front()).(*TreeNode)
		rightNode := queue.Remove(queue.Front()).(*TreeNode)
		if leftNode == nil && rightNode == nil {
			return true
		}
		if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
			return false
		}
		queue.PushBack(leftNode.Left)
		queue.PushBack(rightNode.Right)
		queue.PushBack(leftNode.Right)
		queue.PushBack(rightNode.Left)
	}
	return true
}
