package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	//前序递归，遍历整棵树不带返回值
	var result [][]int
	if root == nil {
		return result
	}
	var path = []int{}
	var travel func(node *TreeNode, sum int, path []int, result *[][]int)
	travel = func(node *TreeNode, sum int, path []int, result *[][]int) {
		//中
		sum -= node.Val //这两句话改写到函数里就是间接回溯,其实也隐藏了回溯
		path = append(path, node.Val)
		//终止条件
		if sum == 0 && node.Left == nil && node.Right == nil {
			*result = append(*result, append([]int{}, path...))
			return
		}
		if node.Left != nil {
			travel(node.Left, sum, path, result)
		}
		if node.Right != nil {
			travel(node.Right, sum, path, result)
		}
		return
	}
	travel(root, targetSum, path, &result)
	return result
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	//前序递归，遍历整棵树不带返回值
	var result [][]int
	if root == nil {
		return result
	}
	var path = []int{}

	var travel func(node *TreeNode, sum int)
	travel = func(node *TreeNode, sum int) {
		//中
		sum -= node.Val //这两句话改写到函数里就是间接回溯,其实也隐藏了回溯
		path = append(path, node.Val)
		//终止条件
		if sum == 0 && node.Left == nil && node.Right == nil {
			result = append(result, append([]int{}, path...))
			return
		}
		if node.Left != nil {
			travel(node.Left, sum)
			path = path[:len(path)-1]

		}
		if node.Right != nil {
			travel(node.Right, sum)
			path = path[:len(path)-1]

		}
		return
	}
	travel(root, targetSum)
	return result
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	//前序递归，遍历整棵树不带返回值
	var result [][]int
	if root == nil {
		return result
	}
	var path = []int{}
	var sum = targetSum
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		//中
		sum -= node.Val //这两句话改写到函数里就是间接回溯,其实也隐藏了回溯!!,错了这就是间接回溯
		path = append(path, node.Val)
		//终止条件
		if sum == 0 && node.Left == nil && node.Right == nil {
			result = append(result, append([]int{}, path...))
			return
		}
		if node.Left != nil {
			travel(node.Left)
			sum += path[len(path)-1]
			path = path[:len(path)-1]

		}
		if node.Right != nil {
			travel(node.Right)
			sum += path[len(path)-1] //sum的回溯写法不能用sum += node.Val
			path = path[:len(path)-1]

		}
		return
	}
	travel(root)
	return result
}

//
func pathSum(root *TreeNode, targetSum int) [][]int {
	//前序递归，遍历整棵树不带返回值
	var result [][]int
	if root == nil {
		return result
	}
	var path = []int{}

	var travel func(node *TreeNode, sum int)
	travel = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		//终止条件
		if sum == node.Val && node.Left == nil && node.Right == nil {
			path = append(path, node.Val)
			result = append(result, append([]int{}, path...))
			path = path[:len(path)-1]
			return
		}
		//中
		sum -= node.Val //这两句话改写到函数里就是间接回溯,其实也隐藏了回溯
		path = append(path, node.Val)
		travel(node.Left, sum)
		travel(node.Right, sum)
		sum += path[len(path)-1]
		path = path[:len(path)-1]
		return
	}
	travel(root, targetSum)
	return result
}
