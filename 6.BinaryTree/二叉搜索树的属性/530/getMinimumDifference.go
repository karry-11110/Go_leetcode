package main

//方法一：中序遍历然后计算最小差值
func getMinimumDifference(root *TreeNode) int {
	var res []int
	findMIn(root, &res)
	min := 1000000 //一个比较大的值
	for i := 1; i < len(res); i++ {
		tempValue := res[i] - res[i-1]
		if tempValue < min {
			min = tempValue
		}
	}
	return min
}

// //中序遍历
func findMIn(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	findMIn(root.Left, res)
	*res = append(*res, root.Val)
	findMIn(root.Right, res)
}

//中序遍历的同时计算最小值
// func getMinimumDifference(root *TreeNode) int {
// 	min := math.MaxInt64
// 	var pre *TreeNode
// 	var travel func(node *TreeNode)
// 	travel = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		travel(node.Left)
// 		if pre != nil && node.Val-pre.Val < min {
// 			min = node.Val - pre.Val
// 		}
// 		pre = node
// 		travel(node.Right)
// 	}
// 	travel(root)
// 	return min

// }
