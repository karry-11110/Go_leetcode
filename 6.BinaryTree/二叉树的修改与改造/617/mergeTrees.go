package main

import "container/list"

//前序递归
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	//终止条件
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	//合并根节点
	root1.Val = root1.Val + root2.Val
	//处理左子树
	root1.Left = mergeTrees(root1.Left, root2.Left)
	//处理右子树
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

//迭代模拟层次遍历,借助两个队列，一个也可以
//这里注意什么节点进队列就可以了，弄清本质，不管什么遍历，只是如何访问它，树的结点还是在那
func mergeTrees_1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	//终止条件
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue1, queue2 := list.New(), list.New()
	queue1.PushBack(root1)
	queue2.PushBack(root2)
	for queue1.Len() > 0 {
		node1 := queue1.Remove(queue1.Front()).(*TreeNode)
		node2 := queue2.Remove(queue2.Front()).(*TreeNode)
		node1.Val += node2.Val
		if node1.Left != nil && node2.Left != nil {
			queue1.PushBack(node1.Left)
			queue2.PushBack(node2.Left)

		}
		if node1.Right != nil && node2.Right != nil {
			queue1.PushBack(node1.Right)
			queue2.PushBack(node2.Right)
		}
		if node1.Left == nil {
			node1.Left = node2.Left
		}
		if node1.Right == nil {
			node1.Right = node2.Right
		}
	}
	return root1
}
