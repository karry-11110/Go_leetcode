package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	//删除节点的五种情况
	//1.没找到，遍历到空节点直接返回
	if root == nil {
		return nil
	}
	if root.Val == key {
		//2.左右孩子都为空
		if root.Left == nil && root.Right == nil {
			return nil
		}
		//3.左孩子为空
		if root.Left == nil {
			return root.Right
		}
		//4.右孩子为空
		if root.Right == nil {
			return root.Left
		}
		//5.左右孩子都有
		newRoot := root.Right //找一个新根节点，右子树的最左边的节点
		for newRoot.Left != nil {
			newRoot = newRoot.Left
		}
		root.Val = newRoot.Val               //把新根节点的值赋给老节点就行了
		root.Right = deleteNode1(root.Right) //删除那个新根节点的空间，因为已经用了它的值
		return root
	}
	//左
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	//右
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	return root
}
func deleteNode1(node *TreeNode) *TreeNode {
	if node.Left == nil {
		tmp := node.Right
		node.Right = nil
		return tmp
	}
	node.Left = deleteNode1(node.Left)
	return node
}
