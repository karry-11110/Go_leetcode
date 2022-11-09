func levelOrder(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		length := st.Len()
		for i := 0; i < length; i++ {
			node := st.Remove(st.Front()).(*TreeNode)
			result = append(result, node.Val)
			if node.Left != nil {
				st.PushBack(node.Left)
			}
			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
	}
	return result
}