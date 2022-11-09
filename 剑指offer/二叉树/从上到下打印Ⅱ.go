func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		subresult := []int{}
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			subresult = append(subresult, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, subresult)
	}
	return result
}