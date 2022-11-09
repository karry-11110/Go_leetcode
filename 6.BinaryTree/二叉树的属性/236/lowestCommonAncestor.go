package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//后序遍历自带回溯
	if root == nil || root == p || root == q {
		return root
	}
	//左
	left := lowestCommonAncestor(root.Left, p, q)
	//右
	right := lowestCommonAncestor(root.Right, p, q)
	//中 回溯过程
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
