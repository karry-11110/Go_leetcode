package main

func isSubTree(root *TreeNode, subRoot *TreeNode) bool {
	//终止条件
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if isSameTree(root, subRoot) == true {
		return true
	}
	//左
	left := isSameTree(root.Left, subRoot)
	//右
	right := isSameTree(root.Right, subRoot)
	//中
	result := left || right
	return result
}
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
