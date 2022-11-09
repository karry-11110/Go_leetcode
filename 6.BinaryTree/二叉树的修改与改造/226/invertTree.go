package main

import (
	"container/list"
)

func invertTree(root *TreeNode) *TreeNode {
	//终止条件
	if root == nil {
		return nil
	}
	//左
	node1 := invertTree(root.Left)
	//右
	node2 := invertTree(root.Right)
	//中
	root.Left = node2
	root.Right = node1
	return root

}

func invertTree_1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}

	}
	return root
}
