//求二叉树深度
func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	st := list.New()
	st.PushBack(pRoot)
	result := 0
	for st.Len() > 0 {
		length := st.Len()
		for i := 0; i < length; i++ {
			node := st.Remove(st.Front()).(*TreeNode)
			if node.Left != nil {
				st.PushBack(node.Left)
			}
			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
		result++
	}
	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxDepth(root *TreeNode) int {
	depth := 0
	result := 0
	var backtracking func(node *TreeNode)
	backtracking = func(node *TreeNode) {
		if node == nil {
			if depth > result {
				result = depth
			}
			return
		}
		depth++
		backtracking(node.Left)
		backtracking(node.Right)
		depth--

		return
	}
	backtracking(root)
	return result
}
