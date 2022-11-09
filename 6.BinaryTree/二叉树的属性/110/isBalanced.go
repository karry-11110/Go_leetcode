package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//左
	left := isBalanced(root.Left)
	//右
	right := isBalanced(root.Right)
	//中
	if left == false || right == false {
		return false
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if abs(leftHeight-rightHeight) > 1 {
		return false
	}
	return true
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func getHeight(root *TreeNode) int {
	//终止条件
	if root == nil {
		return 0
	}
	//左
	maxDepthLeft := getHeight(root.Left)
	//右
	maxDepthRight := getHeight(root.Right)
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
