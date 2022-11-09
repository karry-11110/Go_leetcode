//判断条件的写法根据你的递归东西来写
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var travasal func(left, right *TreeNode) bool
	travasal = func(left, right *TreeNode) bool {
		if right == nil {
			return true
		}
		if left == nil || left.Val != right.Val {
			return false
		}
		return travasal(left.Left, right.Left) && travasal(left.Right, right.Right)
	}
	if A == nil || B == nil {
		return false
	}
	if travasal(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}