package main

import "strconv"

//这里在力扣中函数定义方式的实现有差别！！！！
func binaryTreePaths(root *TreeNode) []string {
	//前序递归，遍历整棵树不带返回值
	var result []string
	if root == nil {
		return result
	}
	var constructPaths func(node *TreeNode, path string, result *[]string)
	constructPaths = func(node *TreeNode, path string, result *[]string) {
		//终止条件
		if node.Left == nil && node.Right == nil {
			paths := path + strconv.Itoa(node.Val)
			*result = append(*result, paths)
			return
		}
		//中
		//左
		if node.Left != nil {
			constructPaths(node.Left, path+strconv.Itoa(node.Val)+"->", result) //间接回溯
		}
		//右
		if node.Right != nil {
			constructPaths(node.Right, path+strconv.Itoa(node.Val)+"->", result)
		}
	}
	constructPaths(root, "", &result)
	return result

}

// func binaryTreePaths(root *TreeNode) []string {
// 	result := make([]string, 0)
// 	var path []int
// 	var travel func(node *TreeNode)
// 	travel = func(node *TreeNode) {
// 		path = append(path, node.Val)
// 		if node.Left == nil && node.Right == nil {

// 			tmp := strconv.Itoa(path[0])
// 			for i := 1; i < len(path); i++ {
// 				tmp = tmp + "->" + strconv.Itoa(path[i])
// 			}
// 			result = append(result, tmp)
// 			return
// 		}

// 		if node.Left != nil {
// 			travel(node.Left)
// 			path = path[:len(path)-1]
// 		}
// 		if node.Right != nil {
// 			travel(node.Right)
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	travel(root)
// 	return result
// }

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	path := []int{}
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			path = append(path, node.Val)
			tmp := strconv.Itoa(path[0])
			for i := 1; i < len(path); i++ {
				tmp = tmp + "->" + strconv.Itoa(path[i])
			}
			result = append(result, tmp)
			path = path[:len(path)-1]
			return
		}
		path = append(path, node.Val)
		travel(node.Left)
		travel(node.Right)
		path = path[:len(path)-1]
	}
	travel(root)
	return result
}
