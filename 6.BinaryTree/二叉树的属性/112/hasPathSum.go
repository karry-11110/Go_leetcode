package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	var result = false
	if root == nil {
		return result
	}
	var pathSum func(node *TreeNode, sum int, targetSum int)
	pathSum = func(node *TreeNode, sum int, targetSum int) {
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			result = true
			return
		}
		if node.Left != nil && result == false {
			pathSum(node.Left, sum, targetSum)
		}
		if node.Right != nil && result == false {
			pathSum(node.Right, sum, targetSum)
		}
	}
	pathSum(root, 0, targetSum)
	return result

}
func hasPathSum_1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var travel func(node *TreeNode, sum int) bool
	travel = func(node *TreeNode, sum int) bool {
		if node.Left == nil && node.Right == nil && sum == 0 {
			return true
		}
		if node.Left == nil && node.Right == nil {
			return false
		}
		if node.Left != nil {
			if travel(node.Left, sum-node.Left.Val) {
				return true
			}
		}
		if node.Right != nil {
			if travel(node.Right, sum-node.Right.Val) {
				return true
			}
		}
		return false
	}
	return travel(root, targetSum-root.Val)

}
func hasPathSum_11(root *TreeNode, targetSum int) bool {
	//前序递归，遍历一条边带返回值，直接返回的写法
	var result bool
	if root == nil {
		return result
	}
	var pathSum func(node *TreeNode, sum int) bool
	pathSum = func(node *TreeNode, sum int) bool {
		//中
		sum -= node.Val //这句话改写到函数里就是间接回溯
		//终止条件
		if node.Left == nil && node.Right == nil && sum == 0 {
			return true
		}

		//左
		if node.Left != nil {
			if pathSum(node.Left, sum) {
				return true
			}
		}
		if node.Right != nil {
			if pathSum(node.Right, sum) {
				return true
			}
		}
		return false
	}
	return pathSum(root, targetSum)
}

func hasPathSum_2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	//左
	left := hasPathSum(root.Left, targetSum-root.Val)
	//右
	right := hasPathSum(root.Right, targetSum-root.Val)
	return left || right
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	targetSum -= root.Val
	left := hasPathSum(root.Left, targetSum)
	right := hasPathSum(root.Right, targetSum)
	targetSum += root.Val
	return left || right

}
