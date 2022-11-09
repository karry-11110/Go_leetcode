package main

func findMode(root *TreeNode) []int {
	result := []int{}
	count := 1
	maxCount := 1
	var pre *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}

		pre = node
		travel(node.Right)
	}
	travel(root)
	return result
}
