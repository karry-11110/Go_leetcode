func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	var backtracking func(left, right int) bool
	backtracking = func(left, right int) bool {
		if right <= left {
			return true
		}
		value := postorder[right]
		index := left
		for index < right && postorder[index] < value {
			index++
		}
		for index2 := index; index2 < right; index2++ {
			if postorder[index2] < value {
				return false
			}
		}
		return backtracking(left, index-1) && backtracking(index, right-1)
	}
	return backtracking(0, len(postorder)-1)
}