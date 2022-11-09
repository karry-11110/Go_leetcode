package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	//中
	root := &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  nil,
		Right: nil,
	}
	//左
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	//右
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return root

}
