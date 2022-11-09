func mirrorTree(root *TreeNode) *TreeNode {
	check(root)
	return root
}
func check(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	check(node.Left)
	check(node.Right)
}