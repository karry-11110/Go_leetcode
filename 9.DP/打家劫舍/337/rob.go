package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	result := robTree(root)
	return max(result[0], result[1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func robTree(node *TreeNode) []int {
	// 注意顺序：0:不偷，1:去偷
	if node == nil {
		return []int{0, 0}
	}
	//后序遍历
	left := robTree(node.Left)
	right := robTree(node.Right)
	//中
	//考虑偷当前跟节点
	robNode := node.Val + left[0] + right[0]
	//考虑偷当前节点的左右孩子
	notRobNode := max(left[0], left[1]) + max(right[0], right[1])
	return []int{notRobNode, robNode}

}
