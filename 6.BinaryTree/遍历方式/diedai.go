package tree

import "container/list"

//前序迭代遍历
func preorder(root *TreeNode) []int {
	var result []int
	st := list.New()
	if root == nil {
		return result
	}
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return result
}

//中序迭代遍历
func inorder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	st := list.New()
	curr := root
	for curr != nil || st.Len() > 0 {
		for curr != nil {
			st.PushBack(curr)
			curr = curr.Left
		}
		curr = st.Remove(st.Back()).(*TreeNode)
		result = append(result, curr.Val)
		curr = curr.Right
	}

	return result
}

//后续迭代遍历
func postorder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		result = append(result, node.Val)
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	reverse(result)
	return result
}
func reverse(result []int) {
	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
}

//层次迭代遍历
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
