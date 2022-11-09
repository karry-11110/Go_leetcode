package main

//递归三部曲
// 1.确定递归函数的参数和返回值
// 再来看返回值，递归函数什么时候需要返回值？什么时候不需要返回值？这里总结如下三点：
// 	如果需要搜索整颗二叉树且不用处理递归返回值，递归函数就不要返回值。（这种情况就是本文下半部分介绍的113.路径总和ii）
// 	如果需要搜索整颗二叉树且需要处理递归返回值，递归函数就需要返回值。 （这种情况我们在236. 二叉树的最近公共祖先 (opens new window)中介绍）
// 	如果要搜索其中一条符合条件的路径，那么递归一定需要返回值，因为遇到符合条件的路径了就要及时返回。
// 2.确定终止条件（防止操作系统的内存栈溢出）
// 3.确定单层递归的逻辑
//前序递归遍历
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return result
}

//中序递归遍历
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		result = append(result, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return result
}

//后续递归遍历
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		result = append(result, node.Val)
	}
	traversal(root)
	return result
}
