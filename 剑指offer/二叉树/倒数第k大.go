func kthLargest(root *TreeNode, k int) int {
	index := k
	result := 0
	var track func(node *TreeNode)
	track = func(node *TreeNode) {
		if node == nil {
			return
		}
		track(node.Right)
		index--
		if index == 0 {
			result = node.Val
			return
		}
		track(node.Left)
	}
	track(root)
	return result
}