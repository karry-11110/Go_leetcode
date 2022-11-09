func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		length := st.Len()
		subresult := []int{}
		for i := 0; i < length; i++ {
			node := st.Remove(st.Front()).(*TreeNode)
			subresult = append(subresult, node.Val)
			if node.Left != nil {
				st.PushBack(node.Left)
			}
			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
		result = append(result, subresult)
	}
	for i, _ := range result {
		if i%2 != 0 {
			reverse(result[i])
		}
	}
	return result
}
func reverse(s []int) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	level := 0
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		length := st.Len()
		subresult := []int{}
		for i := 0; i < length; i++ {
			node := st.Remove(st.Front()).(*TreeNode)
			if level%2 == 0 {
				subresult = append(subresult, node.Val)
			} else {
				subresult = append([]int{node.Val}, subresult...)
			}
			if node.Left != nil {
				st.PushBack(node.Left)
			}
			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
		result = append(result, subresult)
		level++
	}

	return result
}
