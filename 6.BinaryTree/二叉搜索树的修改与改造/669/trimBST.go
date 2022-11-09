package main

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	if root.Val >= low && root.Val <= high {
		root.Left = trimBST(root.Left, low, root.Val)
		root.Right = trimBST(root.Right, root.Val, high)
		return root
	}
	return root
}
