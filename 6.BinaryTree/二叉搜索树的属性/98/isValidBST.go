package main

//二叉搜索树和中序遍历紧密相关，从小到大
// 方法一 ：转换为数组的递归，判断数组是否有序
func isValidBST(root *TreeNode) bool {
	var result []int

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		result = append(result, node.Val)
		traversal(node.Right)
		return
	}
	traversal(root)
	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			return false
		}
	}
	return true
}

// //方法二： 直接递归

// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	return cheak(root, math.MinInt64, math.MaxInt64)
// }
// func cheak(root *TreeNode, min, max int64) bool {
// 	if root == nil {
// 		return true
// 	}
// 	left := cheak(root.Left, min, int64(root.Val))
// 	if min >= int64(root.Val) || max <= int64(root.Val) {
// 		return false
// 	}
// 	right := cheak(root.Right, int64(root.Val), max)
// 	result := left && right
// 	return result
// }

// //中序遍历解法
// func isValidBST(root *TreeNode) bool {
// 	var pre *TreeNode
// 	var travel func(node *TreeNode) bool
// 	travel = func(node *TreeNode) bool {
// 		//截至条件
// 		if node == nil {
// 			return true
// 		}
// 		//左
// 		left := travel(node.Left)
// 		//中
// 		if pre != nil && node.Val <= pre.Val {
// 			return false
// 		}
// 		pre = node
// 		//右
// 		right := travel(node.Right)
// 		return left && right
// 	}
// 	return travel(root)
// }
