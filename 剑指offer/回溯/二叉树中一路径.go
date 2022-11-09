
func pathSum(root *TreeNode, target int) [][]int {
	var result [][]int
	var path []int
	if root == nil {
		return nil
	}
	var backtracking func(node *TreeNode, target int)
	backtracking = func(node *TreeNode, target int) {
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && target == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			backtracking(node.Left, target-node.Left.Val)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			backtracking(node.Right, target-node.Right.Val)
			path = path[:len(path)-1]
		}
		return
	}
	backtracking(root, target-root.Val)
	return result
}