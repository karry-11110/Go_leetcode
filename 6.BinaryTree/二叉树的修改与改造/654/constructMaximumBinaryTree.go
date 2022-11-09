package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	//终止条件
	if len(nums) < 1 {
		return nil
	}
	//从当前数组中找出最大值的下标，即找出根节点
	nowMaxIndex := findMaximumIndex(nums)
	//构造左右子树
	leftTree := nums[:nowMaxIndex]
	rightTree := nums[nowMaxIndex+1:]
	//构造树
	nowRoot := &TreeNode{
		Val:   nums[nowMaxIndex],
		Left:  constructMaximumBinaryTree(leftTree),
		Right: constructMaximumBinaryTree(rightTree),
	}
	return nowRoot
}
func findMaximumIndex(nums []int) int {
	max := -1
	index := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index
}
