package main

//普通二叉树递归
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 1
	//左
	leftNum := countNodes(root.Left)
	//右
	rightNum := countNodes(root.Right)
	//中
	result += leftNum + rightNum
	return result
}

//完全二叉树特性递归
func countNodes_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//左
	leftNum := countNodes(root.Left)
	//右
	rightNum := countNodes(root.Right)
	//中
	left, right := root.Left, root.Right
	leftHeight, rightHeight := 0, 0
	for left != nil {
		left = left.Left
		leftHeight++
	}
	for right != nil {
		right = right.Right
		rightHeight++
	}
	if leftHeight == rightHeight {
		return (2 << leftHeight) - 1
	}
	result := 1 + leftNum + rightNum
	return result
}
